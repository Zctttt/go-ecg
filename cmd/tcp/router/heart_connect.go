package router

import (
	"fmt"
	"go-sv/cmd/tcp/_interface"
	"go-sv/cmd/tcp/instance"
	"go-sv/service/ecg"
)

//HelloZinxRouter Handle
type Connect struct {
	instance.BaseRouter
	Service ecg.Heart
}

func NewYorkConnect() *Connect {
	return &Connect{Service: ecg.NewYork()}
}

func (c *Connect) Handle(request _interface.Request) {
	fmt.Println("Call NewYorkConnect Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
	uuid := c.Service.Connect()
	err := request.GetConnection().SendMsgWithoutPack(uuid[:])
	if err != nil {
		fmt.Println(err)
	}
}
