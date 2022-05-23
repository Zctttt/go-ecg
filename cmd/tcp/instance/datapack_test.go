package instance

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"testing"
)

func TestNewDataPack(t *testing.T) {
	listen, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("listen", err)
	}
	// 模拟服务器
	go func() {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept", err)
		}
		go func(conn net.Conn) {
			// 定义拆包对象
			d := NewDataPack()
			for {

				head := make([]byte, d.GetHeadLen())
				_, err := io.ReadFull(conn, head)
				if err != nil {
					fmt.Println("读头 异常", err)
					return
				}
				msgHead, err := d.UnPack(head)
				if err != nil {
					fmt.Println("解包 异常", err)
				}
				if msgHead.GetMsgLen() > 0 {
					// 当前msg长度>0 开始读取数据
					msg := msgHead.(*Message)
					// 根据数据data长度开辟空间
					msg.Data = make([]byte, msg.DataLen)
					// 读取
					_, err := io.ReadFull(conn, msg.Data)
					if err != nil {
						fmt.Println("读取数据异常", err)
					}
					fmt.Println("数据读取成功")
					fmt.Println("MsgID", strconv.Itoa(int(msg.ID)))
					fmt.Println("DataLen", strconv.Itoa(int(msg.DataLen)))
					fmt.Println("Data", string(msg.Data))
				}
			}
			// 1 读head

			// 2 读内容

		}(conn)
	}()

	// 模拟客户端

	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("客户端创建失败", err)
	}
	d := NewDataPack()

	// 模拟数据粘包 两个数据一起发送
	// 第一个包
	msg1 := &Message{
		ID:      1,
		DataLen: 5,
		Data:    []byte{'h', 'e', 'l', 'l', 'o'},
	}
	msg2 := &Message{
		ID:      2,
		DataLen: 5,
		Data:    []byte{'w', 'o', 'r', 'l', 'd'},
	}
	data1, err := d.Pack(msg1)
	if err != nil {
		fmt.Println("data1 err", err)
	}
	// 第二个包
	data2, err := d.Pack(msg2)
	if err != nil {
		fmt.Println("data2 err", err)
	}
	data := append(data1, data2...)
	//一次性发送
	conn.Write(data)
	select {}
}
