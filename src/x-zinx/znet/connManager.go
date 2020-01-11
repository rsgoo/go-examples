package znet

import (
	"x-zinx/ziface"
	"sync"
	"fmt"
)

type ConnManager struct {
	//管理的连接信息集合
	connections map[uint32]ziface.IConnection

	//保护连接集合的读写锁
	connLock sync.RWMutex
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}

//添加连接
func (connMgr *ConnManager) Add(conn ziface.IConnection) {

	//保护共享资源，加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	connMgr.connections[conn.GetConnID()] = conn

	fmt.Println("connection add to ConnManager successfully：conn num  = ", connMgr.Len())

}

//删除连接
func (connMgr *ConnManager) Remove(conn ziface.IConnection) {
	//保护共享资源，加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()
	delete(connMgr.connections, conn.GetConnID())
	fmt.Println("connection remove to ConnManager successfully：conn num  = ", connMgr.Len())
}

//根据ConnID获取连接
func (connMgr *ConnManager) Get(connID uint32) (conn ziface.IConnection) {
	connMgr.connLock.RLock()
	defer connMgr.connLock.RUnlock()
	conn, ok := connMgr.connections[connID]
	if !ok {
		fmt.Println("connection get connID = ", connID, " not found")
		return nil
	}
	return
}

//获取当前的连接数
func (connMgr *ConnManager) Len() int {
	return len(connMgr.connections)
}

//清除并终止所有的连接
func (connMgr *ConnManager) ClearConn() {
	//保护共享资源，加写锁
	connMgr.connLock.Lock()
	defer connMgr.connLock.Unlock()

	for connID, conn := range connMgr.connections {

		//停止
		conn.Stop()

		//删除
		delete(connMgr.connections, connID)
	}

	fmt.Println("Clear All connections succ! conn num = ", connMgr.Len())
}
