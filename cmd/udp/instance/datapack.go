package instance

import (
	"bytes"
	"encoding/binary"
	"go-sv/cmd/udp/_interface"
)

var defaultHeaderLen int = 4

type DataPack struct {
}

func NewDataPack() _interface.Packet {
	return &DataPack{}
}

func (d *DataPack) GetHeadLen() int {
	// DataLen uint32 4字节  +  ID  uint32 4字节
	return defaultHeaderLen
}

func (d *DataPack) Pack(msg _interface.Message) ([]byte, error) {
	// 创建一个存放bytes的字节缓冲
	buf := bytes.NewBuffer([]byte{})

	// 将 MsgID 写入 buf
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}

	// 将 Data 写入 buf
	if err := binary.Write(buf, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (d *DataPack) UnPack(data []byte) (_interface.Message, error) {
	// 创建一个从输入二进制读取数据的 IOReader
	buf := bytes.NewReader(data)

	msg := &Message{}

	// 读 MsgID
	if err := binary.Read(buf, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}

	return msg, nil

}
