package znet

import (
	"testing"
	"net"
	"fmt"
	"io"
)

func TestDataPack(t *testing.T) {
	/*
	模拟服务器
	 */

	//从客户端读取数据，拆包处理
	//创建socket
	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("server listen err: ", err)
		return
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Accept listen err: ", err)
				return
			}
			go func(conn net.Conn) {
				// 处理客户端的请求
				// 【拆包过程】
				dp := NewDataPack()

				for {

					// 第一次从conn读取，把包的head读取出来
					headData := make([]byte, dp.GetHeadLen())

					//读取conn得数据到headData
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("read head err:", err)
						break
					}
					msgHead, err := dp.Unpack(headData)
					if err != nil {
						fmt.Println("server unpack err:", err)
						return
					}
					// 第二次从conn读取，把包的head中的dataLen读取出来, 在去读取data内容

					if msgHead.GetDataLen() > 0 {
						//msg存在数据，需要第二次读取
						//第二次从conn读取，根据head中的dataLen, 在读取data内容
						msg := msgHead.(*Message)
						//msg := Message{}
						msg.Data = make([]byte, msg.GetDataLen())

						//根据dataLen的长度再次从io流中读取
						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							fmt.Println("server unpack data err:", err)
							return
						}

						//完整的一个消息已经读取完毕
						fmt.Println("--->receive dataLan = ", msg.DataLen)
						fmt.Println("--->receive msgID   = ", msg.ID)
						fmt.Println("--->receive msgData = ", string(msg.Data))
					}

				}

			}(conn)
		}
	}()

	/*
	模拟客户端
	*/
	//模拟一个封包
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial err:", err)
		return
	}

	//创建一个风暴对象 dp
	dp := NewDataPack()

	//模拟粘包过程，封装两个msg 一同发送

	//封装包msg1
	msg1 := &Message{
		ID:      1,
		DataLen: 4,
		Data:    []byte{'z', 'i', 'n', 'x'},
	}
	//打包
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg1 err :", err)
		return
	}
	//封装包msg2
	//封装包msg1
	msg2 := &Message{
		ID:      2,
		DataLen: 6,
		Data:    []byte{'h', 'e', 'l', 'l', ':', ')'},
	}
	//打包
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack msg1 err :", err)
		return
	}

	sendData1 = append(sendData1, sendData2...)

	//发送数据给服务端
	conn.Write(sendData1)

	//客户端阻塞
	select {}
}
