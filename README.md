# xrpc-go

### 同一连接下的多个请求，服务端采用顺序执行的方法，保证响应顺序

### 一个客户端可以创建多个连接
#### 通过服务发现获取到服务所在ip和port，然后请求建立连接，并通过此连接获取服务。