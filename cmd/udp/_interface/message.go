package _interface

//Message 消息抽象层

type Message interface {
	// GetMsgID 获取消息ID
	GetMsgID() uint32
	// GetMsgLen 获取消息长度
	//GetMsgLen() uint32
	// GetData 获取消息内容
	GetData() []byte

	// SetMsgID 设置消息ID
	SetMsgID(uint32)
	// SetMsgLen 设置消息长度
	//SetMsgLen(uint32)
	// SetMsgData 设置消息内容
	SetMsgData([]byte)
}
