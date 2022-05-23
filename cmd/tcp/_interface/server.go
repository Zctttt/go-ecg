package _interface

// Server 服务器抽象层
type Server interface {
	// Start 启动服务器
	Start()
	// Stop 停止服务器
	Stop()
	// Serve 运行服务器
	Serve()

	// AddRouter 路由功能
	AddRouter(msgId uint32, router Router)
	// GetConnMgr 得到链接管理
	GetConnMgr() ConnectionManager
	// SetOnConnStart 设置该Server的连接创建时Hook函数
	SetOnConnStart(func(Connection))
	// SetOnConnStop 设置该Server的连接断开时的Hook函数
	SetOnConnStop(func(Connection))
	// CallOnConnStart 调用连接OnConnStart Hook函数
	CallOnConnStart(conn Connection)
	// CallOnConnStop 调用连接OnConnStop Hook函数
	CallOnConnStop(conn Connection)

	Packet() Packet
}
