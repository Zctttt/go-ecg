package instance

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"go-sv/cmd/tcp/_interface"
	"go-sv/cmd/utils"
)

var defaultHeaderLen uint32 = 8

type DataPack struct {
}

func NewDataPack() _interface.Packet {
	return &DataPack{}
}

func (d *DataPack) GetHeadLen() uint32 {
	// DataLen uint32 4字节  +  ID  uint32 4字节
	return defaultHeaderLen
}

func (d *DataPack) Pack(msg _interface.Message) ([]byte, error) {
	// 创建一个存放bytes的字节缓冲
	buf := bytes.NewBuffer([]byte{})

	// 将 DataLen 写入 buf
	if err := binary.Write(buf, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}

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
	// 读 DataLen
	if err := binary.Read(buf, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	// 读 MsgID
	if err := binary.Read(buf, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}

	// 判断 DataLen 是否超出允许的最大包长度
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New(fmt.Sprintf("当前包长度超出全局预设值  utils.GlobalObject.MaxPackageSize : %d", utils.GlobalObject.MaxPackageSize))
	}

	return msg, nil

}
