package znet

type Message struct {
	//消息的ud
	ID uint32

	//消息长度
	DataLen uint32

	//消息内容
	Data []byte
}

//创建一个Message消息包的方法
func NewMsgPackage(msgID uint32, data []byte) *Message {
	return &Message{
		ID:      msgID,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

//获取 消息id
func (msg *Message) GetMsgID() uint32 {
	return msg.ID
}

//获取消息的长度
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

//获取消息的内容
func (msg *Message) GetData() []byte {
	return msg.Data
}

//设置消息id
func (msg *Message) SetMsgID(id uint32) {
	msg.ID = id
}

//设置消息的长度
func (msg *Message) SetDataLen(len uint32) {
	msg.DataLen = len
}

//设置消息的内容
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}
