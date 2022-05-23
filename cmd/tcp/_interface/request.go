package _interface

// Request 消息抽象层
type Request interface {
	// GetConnection 得到当前连接
	GetConnection() Connection
	// GetData 得到当前请求 Msg 的数据
	GetData() []byte
	// GetMsgID 得到当前请求 Msg 的ID
	GetMsgID() uint32
}
