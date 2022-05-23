package router

import (
	"fmt"
	"go-sv/cmd/udp/_interface"
	"go-sv/cmd/udp/instance"
	"go-sv/service/ecg"
)

//ping test 自定义路由
type Heart struct {
	instance.BaseRouter
	Service ecg.Heart
}

func NewYorkHeart() *Heart {
	return &Heart{Service: ecg.NewYork()}
}

//Ping Handle
func (h *Heart) Handle(request _interface.Request) {

	fmt.Println("心跳数据接入")
	fmt.Println("接收信息 : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
	h.Service.Set(request.GetData())

}
