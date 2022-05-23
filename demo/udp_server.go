package main

import (
	"fmt"
	"net"
	"os"
)

// 限制goroutine数量
var limitChan = make(chan bool, 1000)

// UDP goroutine 实现并发读取UDP数据
func udpProcess(conn *net.UDPConn, i int) {
	fmt.Println(i)
	// 最大读取数据大小
	data := make([]byte, 1024)
	n, add, err := conn.ReadFromUDP(data)
	fmt.Println(n)
	if err != nil {
		fmt.Println("failed read udp msg, error: " + err.Error())
	}
	str := string(data[:n])
	fmt.Println("receive from client, data:" + str)
	_, _ = conn.WriteTo([]byte("hello"), add)
	<-limitChan
}

func udpServer(address string) {
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	conn, err := net.ListenUDP("udp", udpAddr)
	defer conn.Close()
	if err != nil {
		fmt.Println("read from connect failed, err:" + err.Error())
		os.Exit(1)
	}
	i := 0
	for {
		i++
		limitChan <- true
		go udpProcess(conn, i)
	}
}

func main() {
	address := "0.0.0.0:9000"
	udpServer(address)
}
