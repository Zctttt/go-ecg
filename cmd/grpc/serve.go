package grpc

import (
	"context"
	"fmt"
	"go-sv/cmd/protos"
	. "go-sv/service/ecg"
	"google.golang.org/grpc"
	"net"
)

// 重新实现echoServer 重写业务
type heartServer struct {
	protos.UnimplementedHeartServiceServer
}

// 将接收到的请求直接返回
func (h *heartServer) GetUnaryHeart(ctx context.Context, req *protos.HeartRequest) (*protos.HeartResponse, error) {
	uuid := req.GetUUID()
	timeIndex := req.GetTIME()
	data, err := NewYork().Get(uuid, timeIndex)
	if err != nil {
		fmt.Println(err)
		return &protos.HeartResponse{}, nil
	}
	return &protos.HeartResponse{
		UUID:     data.GetUUID(),
		TIME:     data.GetTIME(),
		BPM:      data.GetBPM(),
		HeartADC: data.GetHeartADC(),
	}, nil
}

func Serve() {
	rpcs := grpc.NewServer()
	protos.RegisterHeartServiceServer(rpcs, new(heartServer))
	lis, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	rpcs.Serve(lis)
}
