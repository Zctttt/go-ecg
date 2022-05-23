package instance

type Message struct {
	// 消息ID
	ID uint32
	// 消息长度
	DataLen uint32
	// 消息内容
	Data []byte
}

func NewMessagePackage(id uint32, data []byte) *Message {
	return &Message{
		ID:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

func (m *Message) GetMsgID() uint32 {
	return m.ID
}

func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMsgID(ID uint32) {
	m.ID = ID
}

func (m *Message) SetMsgLen(len uint32) {
	m.DataLen = len
}

func (m *Message) SetMsgData(data []byte) {
	m.Data = data
}
