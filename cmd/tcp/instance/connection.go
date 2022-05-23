package instance

import (
	"context"
	"errors"
	"fmt"
	"go-sv/cmd/tcp/_interface"
	"go-sv/cmd/utils"
	"io"
	"net"
	"sync"
)

type Connection struct {

	//当前Conn属于哪个Server
	TcpServer _interface.Server //当前conn属于哪个server，在conn初始化的时候添加即可

	// 当前连接的socket
	Conn *net.TCPConn

	// 连接的ID
	ConnID uint32

	// 当前的连接状态
	isClosed bool

	// 告知当前连接已经退出/停止的 channel
	ctx    context.Context
	cancel context.CancelFunc

	// 连接路由
	MsgHandler _interface.MsgHandle

	//无缓冲管道，用于读、写两个goroutine之间的消息通信
	msgChan chan []byte

	//有关冲管道，用于读、写两个goroutine之间的消息通信
	msgBuffChan chan []byte
	sync.RWMutex
	//链接属性
	property map[string]interface{}

	//保护链接属性修改的锁
	propertyLock sync.RWMutex
}

// NewConnection 创建一个新的连接
func NewConnection(server _interface.Server, conn *net.TCPConn, connID uint32, msgHandler _interface.MsgHandle) *Connection {
	c := &Connection{
		TcpServer:   server,
		Conn:        conn,
		ConnID:      connID,
		isClosed:    false,
		MsgHandler:  msgHandler,
		msgChan:     make(chan []byte), //msgChan初始化
		msgBuffChan: make(chan []byte, utils.GlobalObject.MaxMsgChanLen),
		property:    make(map[string]interface{}),
	}

	c.GetTCPServer().GetConnMgr().Add(c) //将当前新创建的连接添加到ConnManager中
	return c
}

// SetProperty 设置链接属性
func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	if c.property == nil {
		c.property = make(map[string]interface{})
	}
	c.property[key] = value
}

// GetProperty 获取链接属性
func (c *Connection) GetProperty(key string) (interface{}, error) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()

	if value, ok := c.property[key]; ok {
		return value, nil
	}
	return nil, errors.New("no property found")

}

// RemoveProperty 移除链接属性
func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	delete(c.property, key)
}

// Start 启动连接
func (c *Connection) Start() {
	c.ctx, c.cancel = context.WithCancel(context.Background())
	//1 开启用户从客户端读取数据流程的Goroutine
	go c.startReader()
	//2 开启用于写回客户端数据流程的Goroutine
	go c.startWriter()

	c.TcpServer.CallOnConnStart(c)
	select {
	case <-c.ctx.Done():
		c.finalizer()
		return
	}
}

// Stop 停止连接
func (c *Connection) Stop() {
	c.cancel()
}

// GetTCPConnection 获取当前绑定的socket
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// GetTCPServer  获取当前绑定的socket
func (c *Connection) GetTCPServer() _interface.Server {
	return c.TcpServer
}

// GetConnID 获取当前连接模块的连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// RemoteAddr 获取客户端的TCP状态 IP PORT
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// SendMsg 发送数据
func (c *Connection) SendMsg(msgID uint32, data []byte) error {
	c.RLock()
	defer c.RUnlock()
	if c.isClosed == true {
		return errors.New("发送Msg时连接关闭")
	}
	// 封包
	dp := c.TcpServer.Packet()
	msgPackage, err := dp.Pack(NewMessagePackage(msgID, data))
	if err != nil {
		fmt.Println("封包异常 msgID:", msgID, "err", err)
		return errors.New("封包异常")
	}
	//写回客户端
	c.msgChan <- msgPackage
	return nil
}

func (c *Connection) startReader() {
	fmt.Println("读取开始 。。 ConnID：", c.ConnID, "远程addr", c.RemoteAddr())
	defer fmt.Println("读取结束 。。 ConnID：", c.ConnID, "远程addr", c.RemoteAddr())
	defer c.Stop()
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			// 读取客户端 Msg 的头
			headData := make([]byte, c.TcpServer.Packet().GetHeadLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
				fmt.Println("读头 异常", err)
				return
			}
			// 拆包 得到 MsaID 和 MsgDataLen
			msg, err := c.TcpServer.Packet().UnPack(headData)
			if err != nil {
				fmt.Println("解包 异常", err)
				continue
			}
			var data []byte
			// 根据 MsgDataLen 得到当前包的数据长度
			if msg.GetMsgLen() > 0 {
				data = make([]byte, msg.GetMsgLen())
				//// 当前msg长度>0 开始读取数据
				//msg := msg.(*Message)
				//// 根据数据data长度开辟空间
				//msg.Data = make([]byte, msg.DataLen)
				// 读取
				if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
					fmt.Println("读取数据异常", err)
					continue
				}
				msg.SetMsgData(data)

				// 获取当前对象的request树
				req := Request{
					conn: c,
					msg:  msg,
				}
				if utils.GlobalObject.WorkerPoolSize > 0 {
					//已经启动工作池机制，将消息交给Worker处理
					c.MsgHandler.SendMsgToTaskQueue(&req)
				} else {
					//从绑定好的消息和对应的处理方法中执行对应的Handle方法
					go c.MsgHandler.DoMsgHandler(&req)
				}
			}
		}

	}

}

// StartWriter 写消息Goroutine， 用户将数据发送给客户端
func (c *Connection) startWriter() {

	fmt.Println("[Writer Goroutine is running]")
	defer fmt.Println(c.RemoteAddr().String(), "[conn Writer exit!]")

	for {
		select {
		case data := <-c.msgChan:
			//有数据要写给客户端
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("发送数据异常:, ", err, "写连接退出")
				return
			}
			//针对有缓冲channel需要些的数据处理
		case data, ok := <-c.msgBuffChan:
			if ok {
				//有数据要写给客户端
				if _, err := c.Conn.Write(data); err != nil {
					fmt.Println("发送buff数据异常:, ", err, "写连接退出")
					return
				}
			} else {
				fmt.Println("msgBuffChan 通道关闭")
				break
			}
		case <-c.ctx.Done():
			//conn已经关闭
			return
		}
	}
}

func (c *Connection) SendBuffMsg(msgID uint32, data []byte) error {
	c.RLock()
	defer c.RUnlock()
	if c.isClosed == true {
		return errors.New("Connection closed when send buff msg")
	}

	//将data封包，并且发送
	dp := c.TcpServer.Packet()
	msg, err := dp.Pack(NewMessagePackage(msgID, data))
	if err != nil {
		fmt.Println("解包异常  msg ID = ", msgID)
		return errors.New("解包异常")
	}

	//写回客户端
	c.msgBuffChan <- msg

	return nil
}

func (c *Connection) SendMsgWithoutPack(data []byte) error {
	c.RLock()
	defer c.RUnlock()
	if c.isClosed == true {
		return errors.New("发送Msg时连接关闭")
	}
	//写回客户端
	c.msgChan <- data
	return nil
}

// Context 返回ctx，用于用户自定义的go程获取连接退出状态
func (c *Connection) Context() context.Context {
	return c.ctx
}

func (c *Connection) finalizer() {
	//如果用户注册了该链接的关闭回调业务，那么在此刻应该显示调用
	c.TcpServer.CallOnConnStop(c)

	c.Lock()
	defer c.Unlock()

	//如果当前链接已经关闭
	if c.isClosed == true {
		return
	}

	fmt.Println("连接退出 ConnID = ", c.ConnID)

	// 关闭socket链接
	_ = c.Conn.Close()

	//将链接从连接管理器中删除
	c.TcpServer.GetConnMgr().Remove(c)

	//关闭该链接全部管道
	close(c.msgBuffChan)
	//设置标志位
	c.isClosed = true
}
