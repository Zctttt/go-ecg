// @Title
// @Description
// @Author
// @Update
package main

import (
	"context"
	"fmt"
	"go-sv/cmd/protos"
	"google.golang.org/grpc"
	"log"
)

func main() {
	//创建tcp连接8
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// 创建客户端
	client := protos.NewHeartServiceClient(conn)
	i := 0
	//for {
	// 触发rpc请求
	resp, err := client.GetUnaryHeart(context.Background(), &protos.HeartRequest{UUID: "1234", TIME: "5678"})
	if err != nil {
		log.Fatal(err)
	}
	i++
	// 打印返回值
	fmt.Println(resp)

	//}
	return
}
