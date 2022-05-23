package _interface

type Room interface {
	//// Start 开启房间
	//Start()
	//// Stop 关闭房间
	//Stop()
	//// 更新当前房间的数据
	//Flash(ConnID uint32)
	// GetName 获取当前房间名
	GetName() string
	// GetName 设置当前房间名
	SetName(name string) error
	// ResetName 修改当前房间名
	ResetName(name string) error

	Publish(msg []byte)
	SubScribe(connID uint32) (err error)
	Start(roomMsgID uint32)
	Stop()
	//// GetName 获取当前房间名
	//GetID()uint32
	//// GetMasterConnectionID 获取当前房间的所有连接
	//GetMasterConnectionID()uint32
	//// GetConnectionIDs 获取当前房间的所有连接
	//GetConnectionIDs()[]uint32
	//// JoinConnection 加入房间
	//JoinRoom(ConnID uint32)
	//// JoinConnection 加入房间
	//LeaveRoom(ConnID uint32)
	//// SetRoomCreated 设置该Server的连接创建时Hook函数
	//SetRoomCreated(HandlerFunc)
	//// SetRoomDestroyed 设置该Server的连接断开时的Hook函数
	//SetRoomDestroyed(HandlerFunc)
}
