package _interface

// MsgHandle 消息管理抽象层
type MsgHandle interface {
	// DoMsgHandler 马上以非阻塞方式处理消息
	DoMsgHandler(request Request)
	// AddRouter 为消息添加具体的处理逻辑
	AddRouter(msgId uint32, router Router)
	// StartWorkerPool 启动worker工作池
	StartWorkerPool()
	// SendMsgToTaskQueue 将消息交给TaskQueue,由worker进行处理
	SendMsgToTaskQueue(request Request)
}
