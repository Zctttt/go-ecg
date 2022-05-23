// @Title
// @Description
// @Author
// @Update

package main

import (
	"context"
	"fmt"
	"go-sv/protos"
	"net"

	"google.golang.org/grpc"
)

// 重新实现echoServer 重写业务
type echoServer struct {
	protos.UnimplementedEchoServiceServer
}

// 将接收到的请求直接返回
func (e *echoServer) GetUnaryEcho(ctx context.Context, req *protos.EchoRequest) (*protos.EchoResponse, error) {
	res := "recived:" + req.GetReq()
	fmt.Println(res)
	return &protos.EchoResponse{Res: res}, nil
}

func main() {
	rpcs := grpc.NewServer()
	protos.RegisterEchoServiceServer(rpcs, new(echoServer))
	lis, err := net.Listen("tcp", "0.0.0.0:10000")
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	rpcs.Serve(lis)
}
