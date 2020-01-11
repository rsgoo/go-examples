package ziface

//连接管理模块抽象层

type IConnManager interface {
	//添加连接
	Add(conn IConnection)

	//删除连接
	Remove(conn IConnection)

	//根据ConnID获取连接
	Get(connID uint32) (conn IConnection)

	//获取当前的连接数
	Len() int

	//清除并终止所有的连接
	ClearConn()
}
