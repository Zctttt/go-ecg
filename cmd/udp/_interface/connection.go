package _interface

import (
	"context"
	"net"
)

// Connection 连接模块抽象层
type Connection interface {
	// Start 启动连接
	Start()
	// Stop 停止连接
	Stop()
	// Context 返回ctx，用于用户自定义的go程获取连接退出状态
	Context() context.Context

	// GetUDPConnection 获取当前绑定的socket
	GetUDPConnection() *net.UDPConn
	// GetConnID 获取当前连接模块的连接ID
	GetConnID() uint32
	// RemoteAddr 获取客户端的TCP状态 IP PORT
	RemoteAddr() net.Addr

	// SendMsg   发送数据
	SendMsg(msgID uint32, data []byte) error
	// SendBuffMsg 直接将Message数据发送给远程的TCP客户端(有缓冲)
	SendBuffMsg(msgId uint32, data []byte) error //添加带缓冲发送消息接口

	// SetProperty 设置链接属性
	SetProperty(key string, value interface{})
	// GetProperty 获取链接属性
	GetProperty(key string) (interface{}, error)
	// RemoveProperty 移除链接属性
	RemoveProperty(key string)
}

type HandlerFunc func(*net.UDPConn, []byte, int) error
