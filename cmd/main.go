package main

import (
	"go-sv/cmd/grpc"
	Tcp "go-sv/cmd/tcp/instance"
	TcpRouter "go-sv/cmd/tcp/router"
	Udp "go-sv/cmd/udp/instance"
	UdpRouter "go-sv/cmd/udp/router"
)

//
//// DoConnectionBegin 创建连接的时候执行
//func DoConnectionBegin(conn UdpI.Connection) {
//	//log.Debug("DoConnecionBegin is Called ... ")
//
//	//设置两个链接属性，在连接创建之后
//	//log.Debug("Set conn Name, Home done!")
//	//conn.SetProperty("Name", "Aceld")
//	//conn.SetProperty("Home", "https://www.kancloud.cn/@aceld")
//
//	//err := conn.SendMsg(conn., []byte("DoConnection BEGIN..."))
//	//if err != nil {
//	//	log.Error(err)
//	//}
//
//}
//
//// DoConnectionLost 连接断开的时候执行
//func DoConnectionLost(conn UdpI.Connection) {
//	//在连接销毁之前，查询conn的Name，Home属性
//	//if name, err := conn.GetProperty("Name"); err == nil {
//	//	log.Error("Conn Property Name = ", name)
//	//}
//	//
//	//if home, err := conn.GetProperty("Home"); err == nil {
//	//	log.Error("Conn Property Home = ", home)
//	//}
//
//	//log.Debug("DoConneciotnLost is Called ... ")
//}

func main() {
	tcpClient := Tcp.NewServer()
	tcpClient.GetConnMgr()
	tcpClient.AddRouter(1, TcpRouter.NewYorkConnect())
	tcpClient.AddRouter(2, TcpRouter.NewYorkUserList())
	go tcpClient.Serve()

	udpClient := Udp.NewServer()
	udpClient.GetConnMgr()
	udpClient.AddRouter(1684956531, UdpRouter.NewYorkHeart())
	go udpClient.Serve()

	go grpc.Serve()
	select {}
}
