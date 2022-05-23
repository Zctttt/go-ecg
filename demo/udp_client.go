package main

import (
	"fmt"
	"net"
	"os"
	//  "io"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:9000")
	defer conn.Close()
	if err != nil {
		os.Exit(1)
	}
	//dp := znet.NewDataPack()
	//msg, _ := dp.Pack(znet.NewMessagePackage(0, []byte("1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890xx")))
	//conn.Write(msg)
	conn.Write([]byte("    12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678"))
	fmt.Println("send msg")

}
