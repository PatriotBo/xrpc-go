package proto

import "github.com/golang/protobuf/proto"

func UnMarshal(data []byte, v interface{}) error {
	protoMsg := v.(proto.Message)
	protoMsg.Reset()

	if pu, ok := protoMsg.(proto.Unmarshaler); ok {
		return pu.Unmarshal(data)
	}
	return nil
}
