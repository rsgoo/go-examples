package znet

import (
	"x-zinx/ziface"
	"bytes"
	"encoding/binary"
	"x-zinx/utils"
	"errors"
)

// 封包 | 拆包模块
// 直接面向TCP连接中的数据流，用于处理TCP粘包问题

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

//获取包长度方法
func (dp *DataPack) GetHeadLen() uint32 {
	//dataLen uint32(4个字节) + msgID uint32(4个字节)
	return 8
}

//封包
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	//创建一个存放bytes的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	//将dataLen写入到dataBuff中
	err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen())
	if err != nil {
		return nil, err
	}
	//将msgID 写入到 dataBuff中
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgID())
	if err != nil {
		return nil, err
	}

	//将data 写入到 dataBuff中
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetData())
	if err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil

}

//拆包
func (dp *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	//创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)

	msg := &Message{}

	//只解压Head信息，先后得到dataLen和MsgID

	//读取dataLen
	err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen)
	if err != nil {
		return nil, err
	}

	//读取数据类型
	err = binary.Read(dataBuff, binary.LittleEndian, &msg.ID)
	if err != nil {
		return nil, err
	}

	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New(" too large msg data received")
	}

	return msg, nil
}
