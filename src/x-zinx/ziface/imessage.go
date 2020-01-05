package ziface

//将请求的消息封装到一个Message中，定义抽象的接口
type IMessage interface {
	//获取 消息id
	GetMsgID() uint32

	//获取消息的长度
	GetDataLen() uint32

	//获取消息的内容
	GetData() []byte

	//设置消息id
	SetMsgID(uint32)

	//设置消息的长度
	SetDataLen(uint32)

	//设置消息的内容
	SetData([]byte)
}
