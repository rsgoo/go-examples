package znet

import (
	"net"
	"x-zinx/ziface"
	"fmt"
	"io"
	"errors"
	"x-zinx/utils"
	"sync"
)

//连接模块
type Connection struct {
	TCPServer ziface.IServer

	//当前连接的socket TCP套接字
	Conn *net.TCPConn

	//连接id
	ConnID uint32

	//当前的连接状态
	isClosed bool

	//当前连接锁绑定的处理业务的api
	//handleAPI ziface.HandleFun

	//告知当前的连接是否已经退出或是停止,由reader告知writer退出
	exitChan chan bool

	//无缓冲的管道，用于读,写Goroutine之间的消息通信
	msgChan chan []byte

	//该连接处理的方法router
	//Router ziface.IRouter

	//消息管理msgID
	MsgHandler ziface.IMsgHandler

	//连接属性集合
	property map[string]interface{}

	//保护链接的锁
	propertyLock sync.RWMutex
}

//初始化连接模块的方法
func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandler) *Connection {
	c := &Connection{
		TCPServer:  server,
		Conn:       conn,
		ConnID:     connID,
		MsgHandler: msgHandler,
		isClosed:   false,
		msgChan:    make(chan []byte),
		exitChan:   make(chan bool, 1),
		property:   make(map[string]interface{}),
	}

	//将当前的连接加入到connMgr中
	c.TCPServer.GetConnMgr().Add(c)

	return c
}

//连接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("【Reader Goroutine is running...】")

	defer fmt.Println("connID = ", c.ConnID, "【reader is exit】, remote addr is ", c.RemoteArr().String())
	defer c.Stop()

	for {
		//创建一个拆包解包的对象
		dp := NewDataPack()

		//读取客户端的msgHead 二进制流，8个字节
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head err:", err)
			break
		}

		//拆包，得到msgID 和 msgDataLen, 放在msg消息中
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error:", err)
			break
		}

		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error ", err)
				break
			}
		}

		msg.SetData(data)

		//根据dataLen, 再次读取Data，放在msg.data中

		//得到当前conn数据的request请求数据
		req := Request{
			conn: c,
			msg:  msg,
		}

		//从路由中找到注册绑定Router对应的router调用
		//go c.MsgHandler.DoMsgHandler(&req)
		if utils.GlobalObject.WorkerPoolSize > 0 {
			//已经启动工作池机制，将消息交给Worker处理
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			//从绑定好的消息和对应的处理方法中执行对应的Handle方法
			go c.MsgHandler.DoMsgHandler(&req)
		}
	}
}

//写消息Goroutine, 专门发送给客户端消息的模块
func (c *Connection) StartWriter() {
	fmt.Println("【Writer Goroutine is runing】")
	defer fmt.Println(c.RemoteArr().String(), "【conn writer exit!】")
	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send data error:", err)
				return
			}
		case <-c.exitChan:
			//代表Reader已经退出，此时Writer也要退出
			return
		}
	}
}

//启动连接，让当前的连接主备开始工作
func (c *Connection) Start() {
	fmt.Println("Connection start()... ConnID = ", c.ConnID)

	//启动从当前连接读取数据的业务
	go c.StartReader()

	//启动从当前连接写数据的业务
	go c.StartWriter()

	//按照开发者串第进来的，创建连接之后需要调用的函数，执行对应的hook函数
	c.TCPServer.CallOnConnStart(c)

}

//停止连接，结束当前连接的工作
func (c *Connection) Stop() {
	fmt.Println("Connection stop()... ConnID = ", c.ConnID)

	//如果当前连接已经关闭
	if c.isClosed == true {
		return
	}

	c.isClosed = true

	//关闭socket连接
	c.Conn.Close()

	//告知writer关闭
	c.exitChan <- true

	//调用开发者注册的销毁连接之前需要执行的业务Hook函数
	c.TCPServer.CallOnConnStop(c)

	c.TCPServer.GetConnMgr().ClearConn()

	close(c.exitChan)
	close(c.msgChan)
}

//获取当前连接绑定的socket connect
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//获取当前连接模块的连接id
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//获取远程客户端的 tcp 状态 IP port
func (c *Connection) RemoteArr() net.Addr {
	return c.Conn.RemoteAddr()
}

//提供一个sendMsg方法，
func (c *Connection) SendMsg(msgID uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("Connection closed when send msg")
	}

	//将data进行封包
	dp := NewDataPack()
	binaryMsg, err := dp.Pack(NewMsgPackage(msgID, data))
	if err != nil {
		fmt.Println("pack error msg id:", msgID)
		return errors.New("pack error msg")
	}

	//将数据发送给客户端
	//if _, err := c.Conn.Write(binaryMsg); err != nil {
	//	fmt.Println("Write msg id:", msgID, " error:", err)
	//	return errors.New("conn write error")
	//}

	c.msgChan <- binaryMsg
	return nil
}

//设置连接属性
func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	c.property[key] = value
}

//获取连接属性
func (c *Connection) GetProperty(key string) (interface{}, error) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()
	if value, ok := c.property[key]; ok {
		return value, nil
	}
	return nil, errors.New("No Property found")
}

//移除连接属性
func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()
	delete(c.property, key)
}
