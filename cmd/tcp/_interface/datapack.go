package _interface

// 协议抽象层 解决tcp粘包问题
type DataPack interface {
	// GetHeadLen 获取包头的长度方法
	GetHeadLen() uint32
	// Pack 封包方法
	Pack(msg Message) ([]byte, error)
	// UnPack 拆包方法
	UnPack([]byte) (Message, error)
}
