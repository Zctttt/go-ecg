package instance

import "go-sv/cmd/udp/_interface"

type Option func(s *Server)

// WithPacket 只要实现Packet 接口可自由实现数据包解析格式，如果没有则使用默认解析格式
func WithPacket(pack _interface.Packet) Option {
	return func(s *Server) {
		s.packet = pack
	}
}
