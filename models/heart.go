package models

import "github.com/vmihailenco/msgpack"

type HeartData struct {
	HeartADC []byte
	UUID     []byte
	TIME     []byte
	BPM      []byte
}

func NewHeartData(buf []byte) *HeartData {
	return &HeartData{
		UUID:     buf[0:4],
		TIME:     buf[4:8],
		BPM:      buf[8:12],
		HeartADC: buf[12:112],
	}
}

func (s *HeartData) MarshalBinary() ([]byte, error) {
	return msgpack.Marshal(s)
}

func (s *HeartData) UnmarshalBinary(data []byte) error {
	return msgpack.Unmarshal(data, s)
}

func (s *HeartData) GetUUID() string {
	return string(s.UUID)
}

func (s *HeartData) GetTIME() string {
	return string(s.TIME)
}

func (s *HeartData) GetBPM() string {
	return string(s.BPM)
}

func (s *HeartData) GetHeartADC() string {
	return string(s.HeartADC)
}
