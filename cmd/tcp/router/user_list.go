package router

import (
	"fmt"
	"go-sv/cmd/tcp/_interface"
	"go-sv/cmd/tcp/instance"
	"go-sv/service/ecg"
)

//HelloZinxRouter Handle
type UserList struct {
	instance.BaseRouter
	Service ecg.Heart
}

func NewYorkUserList() *UserList {
	return &UserList{Service: ecg.NewYork()}
}

func (c *UserList) Handle(request _interface.Request) {
	fmt.Println("Call NewYorkConnect Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
	err := request.GetConnection().SendMsgWithoutPack(c.Service.UserList())
	if err != nil {
		fmt.Println(err)
	}
}
