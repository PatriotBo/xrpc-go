package xrpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"strings"
	"time"
)

type ZKRtr struct {
	conn          *zk.Conn
	ServiceAll    map[string]*ServiceInfo // 服务信息列表，将zk数据存到本地，并通过监听事件更新数据
	watchSercices map[string]bool         // 监听的服务列表，开始监听即为true
}

//存储在zk中的服务的信息
type ServiceInfo struct {
	Name    string   `json:"name"`     // 服务名称
	IP      string   `json:"ip"`       // ip
	Port    int      `json:"port"`     // 端口
	Weight  int      `json:"weight"`   // 权重
	funcs   []string `json:"funcs"`    // 服务提供的方法名
	SvrType int      `json:"svr_type"` // 服务类型
	SvrId   int      `json:"svr_id"`   // 服务id
	Version int32    `json:"version"`  // 服务版本，每次修改后需要加1
}

/*
	zk上的存储路径为 /rpc/
*/
var rootPath = "/rpc/"

// 创建一个新的zk会话连接，用于注册服务信息
func NewRegister() (ZKRtr, error) {
	var zkr ZKRtr
	var hosts = []string{"127.0.0.1:2181"}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		log.Println("zk connect failed:", err)
		return zkr, err
	}
	zkr.conn = conn
	conn.State()
	zkr.ServiceAll = make(map[string]*ServiceInfo)
	zkr.watchSercices = make(map[string]bool)
	return zkr, nil
}

// 注册服务信息到zk
func (zr *ZKRtr) Register(info ServiceInfo) error {
	flag := int32(0)
	acl := zk.WorldACL(zk.PermAll) // 权限控制-所有权限
	fmt.Println("acl: ", acl)
	data, err := json.Marshal(info)
	if err != nil {
		log.Println("register marshal failed:", err)
		return err
	}
	// 判断节点是否存在，不存在则创建，存在则修改信息
	znodePath := rootPath + info.Name
	isExists, _, err := zr.conn.Exists(znodePath)
	if err != nil {
		log.Println("exists err:", err)
		return err
	}
	if !isExists {
		path, err := zr.conn.Create(znodePath, data, flag, acl)
		if err != nil {
			log.Println("register failed:", err)
			return err
		}
		log.Println("register success path: ", path)
	} else {
		log.Println("znode already exist")
		_, err = zr.conn.Set(znodePath, data, info.Version)
		if err != nil {
			log.Println("znode modify failed:", err.Error())
			return err
		}
	}
	return nil
}

// 从zk中获取服务信息
func (zr *ZKRtr) GetServiceInfoByName(name string) (ServiceInfo, error) {
	if svrInfo, ok := zr.ServiceAll[name]; ok {
		return svrInfo, nil
	}

	var ret ServiceInfo
	znodePath := rootPath + name
	isExists, _, err := zr.conn.Exists(znodePath)
	if err != nil {
		log.Println("get info err：", err)
		return ret, err
	}
	if !isExists {
		log.Printf("znode %s not exists\n", znodePath)
		return ret, errors.New(fmt.Sprintf("znode %s not exists", znodePath))
	}

	data, _, err := zr.conn.Get(znodePath)
	if err != nil {
		log.Println("get info failed:", err)
		return ret, err
	}
	err = json.Unmarshal(data, &ret)
	if err != nil {
		log.Println("unmarshal info failed:", err)
		return ret, err
	}

	// 添加数据到内存中
	zr.ServiceAll[name] = ret

	// 判断是否注册了监听事件，没有则进行注册
	if _, ok := zr.watchSercices[znodePath]; !ok {
		zr.watchSercices[znodePath] = true
		go zr.watchNodeDataChange(znodePath)
	}
	return ret, nil
}

//监听服务watch 监听节点数据和子节点的增删变化 循环监听，因为zk的监听机制是一次性的
func (zr *ZKRtr) watchNodeDataChange(path string) {
	for {
		_, _, ch, err := zr.conn.ChildrenW(path)
		if err != nil {
			log.Println("watch node error")
		}
		e := <-ch
		go zr.watchCallback(e, path)
	}
}

//监听事件回调
func (zr *ZKRtr) watchCallback(event zk.Event, path string) {
	log.Println("watch callback path:", path)
	serviceName := getNameFromPath(path)
	if len(serviceName) <= 0 {
		return
	}
	switch event.Type {
	case zk.EventNodeDeleted: // 节点删除，则取消监听
		delete(zr.ServiceAll, serviceName)
		zr.watchSercices[serviceName] = false

	case zk.EventNodeDataChanged:
		fallthrough
	case zk.EventNodeChildrenChanged: // 数据变化，修改本地内存
		var svrInfo ServiceInfo
		data, _, err := zr.conn.Get(path)
		if err != nil {
			log.Println("get info failed:", err)
			return
		}
		err = json.Unmarshal(data, &svrInfo)
		if err != nil {
			log.Println("unmarshal info failed:", err)
			return
		}
		zr.ServiceAll[serviceName] = &svrInfo
	default:

	}
}

func getNameFromPath(path string) string {
	index := strings.LastIndex(path, "/")
	if index+1 < len(path) {
		return path[index+1:]
	}
	return ""
}
