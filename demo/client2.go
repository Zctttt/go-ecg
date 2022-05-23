package main

import (
	"fmt"
	znet "go-sv/cmd/tcp/instance"
	"io"
	"net"
	"time"
)

/*
   模拟客户端
*/
func main() {

	fmt.Println("Client Test ... start")
	//3秒之后发起测试请求，给服务端开启服务的机会

	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	//for {
	//发封包message消息
	dp := znet.NewDataPack()
	msg, _ := dp.Pack(znet.NewMessagePackage(2, []byte("Zinx V0.6 Client1 Test Message")))
	_, err = conn.Write(msg)
	if err != nil {
		fmt.Println("write error err ", err)
		return
	}

	//先读出流中的head部分
	headData := make([]byte, 4)

	_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
	fmt.Println(string(headData))

	if err != nil {
		fmt.Println("read head error")
		//break
	}
	//将headData字节流 拆包到msg中
	d := string(headData)
	fmt.Println(d)
	select {}
	//msgHead, err := dp.UnPack(headData)
	//if err != nil {
	//	fmt.Println("server unpack err:", err)
	//	return
	//}
	//
	//if msgHead.GetMsgLen() > 0 {
	//	//msg 是有data数据的，需要再次读取data数据
	//	msg := msgHead.(*znet.Message)
	//	msg.Data = make([]byte, msg.GetMsgLen())
	//
	//	//根据dataLen从io中读取字节流
	//	_, err := io.ReadFull(conn, msg.Data)
	//	if err != nil {
	//		fmt.Println("server unpack data err:", err)
	//		return
	//	}
	//
	//	fmt.Println("==> Recv Msg: ID=", msg.ID, ", len=", msg.DataLen, ", data=", string(msg.Data))
	//}

	time.Sleep(1 * time.Second)
	//}
}
