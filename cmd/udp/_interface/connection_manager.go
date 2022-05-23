package _interface

type ConnectionManager interface {
	Add(conn Connection)                   //添加链接
	Remove(conn Connection)                //删除连接
	Get(connID uint32) (Connection, error) //利用ConnID获取链接
	Len() int                              //获取当前连接
	ClearConn()                            //删除并停止所有链接
}
