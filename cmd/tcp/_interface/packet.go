package _interface

type Packet interface {
	UnPack(binaryData []byte) (Message, error)
	Pack(msg Message) ([]byte, error)
	GetHeadLen() uint32
}
