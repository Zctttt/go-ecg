package _interface

type RoomManager interface {
	RoomExist(roomName string) bool
	Add(room Room)                   //添加链接
	Remove(room Room)                //删除连接
	Get(roomID string) (Room, error) //利用ConnID获取链接
	Len() int                        //获取当前房间数
	List() []string                  //获取房间id数组
	ClearRoom()                      //删除并停止所有房间
}
