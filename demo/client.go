package main

import (
	"bytes"
	"fmt"
	znet "go-sv/cmd/tcp/instance"
	"net"
	"time"

	"github.com/aler9/gomavlib/pkg/dialect"
	"github.com/aler9/gomavlib/pkg/dialects/ardupilotmega"
	"github.com/aler9/gomavlib/pkg/parser"
)

/*
   模拟客户端
*/
func main() {

	fmt.Println("Client Test ... start")
	Buf := bytes.NewBuffer(nil)
	dialectDE, err := dialect.NewDecEncoder(ardupilotmega.Dialect)
	if err != nil {
		panic(err)
	}
	writer, err := parser.NewWriter(parser.WriterConf{
		Writer:      Buf,
		DialectDE:   dialectDE,
		OutVersion:  parser.V2, // change to V1 if you're unable to communicate with the target
		OutSystemID: 10,
	})
	//发封包message消息
	dp := znet.NewDataPack()
	msg, _ := dp.Pack(znet.NewMessagePackage(0, Buf.Bytes()))
	fmt.Println(msg)
	//3秒之后发起测试请求，给服务端开启服务的机会

	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		err = writer.WriteMessage(&ardupilotmega.MessageParamValue{
			ParamId:    "first",
			ParamValue: 123456,
			ParamType:  ardupilotmega.MAV_PARAM_TYPE_UINT32,
		})
		//err = writer.WriteMessage(&ardupilotmega.MessageParamValue{
		//	ParamId:    "first",
		//	ParamValue: 123456,
		//	ParamType:  ardupilotmega.MAV_PARAM_TYPE_UINT32,
		//})
		//发封包message消息
		dp := znet.NewDataPack()
		msg, _ := dp.Pack(znet.NewMessagePackage(0, Buf.Bytes()))
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}

		////先读出流中的head部分
		//headData := make([]byte, dp.GetHeadLen())
		//_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
		//if err != nil {
		//	fmt.Println("read head error")
		//	break
		//}
		////将headData字节流 拆包到msg中
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
	}
}
