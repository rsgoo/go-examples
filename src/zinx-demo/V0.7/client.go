package main

import (
	"fmt"
	"time"
	"net"
	"x-zinx/znet"
	"io"
)

//模拟客户端
func main() {
	fmt.Println("client start....")
	time.Sleep(time.Second)

	//创建tcp连接远程服务器，得到一个conn
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}

	for {

		//发送封包的msg消息
		dp := znet.NewDataPack()
		binaryMsg, err := dp.Pack(znet.NewMsgPackage(0, []byte("Zinx  client Test Message")))

		if err != nil {
			fmt.Println("client  pack err :", err)
			return
		}

		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("conn write err: ", err)
			return
		}

		//服务器响应数据message，

		//先读取流中的head部分，得到ID 和 dataLen
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head error:", err)
			return
		}

		//在根据dataLen进行二次读取，将data读出来
		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("client unpack msgHead error:", err)
			break
		}

		if msgHead.GetDataLen() > 0 {
			msg := msgHead.(*znet.Message)

			msg.Data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("read msg data err:", err)
				return
			}
			fmt.Println("receive server msg ID =", msg.ID)
			fmt.Println("receive server msg len =", msg.DataLen)
			fmt.Println("receive server msg data =", string(msg.Data))
			fmt.Println("-----------------------------")
		}

		time.Sleep(time.Second)
	}
}
