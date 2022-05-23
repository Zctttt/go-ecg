package instance

import "go-sv/cmd/udp/_interface"

type Request struct {
	// 已经和客户端建立好的连接
	conn _interface.Connection
	// 客户端请求的数据
	msg _interface.Message
}

func (r *Request) GetConnection() _interface.Connection {
	return r.conn
}
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgID()
}
