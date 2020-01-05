package ziface

//定义个服务器接口
type IServer interface {
	//启动服务器
	Start()

	//运行服务器
	Stop()

	//运行服务器
	Serve()

	//路由功能: 给当前的服务注册一个路由方法，共客户端的连接处理使用
	//AddRouter(router IRouter)
	AddRouter(msgID uint32, router IRouter)
}
