package instance

import (
	"fmt"
	"go-sv/cmd/udp/_interface"
	"go-sv/cmd/utils"
	"net"
)

type Server struct {
	// 连接名
	Name string
	// 连接版本
	IPVersion string
	// 服务器IP
	IP string
	// 服务器端口
	Port int
	//当前Server的消息管理模块，用来绑定MsgId和对应的处理方法
	msgHandler _interface.MsgHandle
	// 当前Server的连接管理器
	ConnMrg _interface.ConnectionManager

	//该Server的连接创建时Hook函数
	OnConnStart func(conn _interface.Connection)
	//该Server的连接断开时的Hook函数
	OnConnStop func(conn _interface.Connection)

	packet _interface.Packet

	exitChan chan bool
}

func NewServer(opts ...Option) _interface.Server {
	utils.GlobalObject.Reload()
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "udp",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.Port,
		msgHandler: NewMsgHandle(),
		ConnMrg:    CoonMrg,
		packet:     NewDataPack(),
		exitChan:   make(chan bool),
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *Server) Start() {
	fmt.Printf("【基本配置】： Name:%s ,IP:%s ,Port:%d", utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.Port)
	fmt.Printf("【基本连接配置】： Version:%s ,MaxConn:%d ,MaxPackageSize:%d", utils.GlobalObject.Version, utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPackageSize)
	fmt.Println("【服务开启】  IP :", s.IP, "   Port:", s.Port)
	go func() {
		s.msgHandler.StartWorkerPool()
		// 1 获取tcp的addr
		addr, err := net.ResolveUDPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("获取tcp addr 异常", err)
			return
		}
		// 2 监听服务器地址
		conn, err := net.ListenUDP(s.IPVersion, addr)
		if err != nil {
			fmt.Println(err)
		}
		defer conn.Close()
		if err != nil {
			fmt.Println("监听 listen 异常", err, "\nIPV:", s.IPVersion, "\nIP", s.IP, "\nPORT", s.Port)
			return
		}
		fmt.Println("【启动成功】，", s.Name, "开启监听")
		var cid uint32
		cid = 0
		// 3 阻塞等待客户端连接，处理客户端连接业务（读写）
		for {
			// 获取连接后 执行之后的动作

			//3.2 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
			if s.ConnMrg.Len() < utils.GlobalObject.MaxConn {
				//执行连接任务
				dealConn := NewConnection(s, conn, cid, s.msgHandler)
				cid++
				//3.4 启动当前链接的处理业务

				go dealConn.Start()
			}
		}

	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name ", s.Name)

	//将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
	s.ConnMrg.ClearConn()
	s.exitChan <- true
}

func (s *Server) Serve() {
	s.Start()
	exitFlag := <-s.exitChan
	if exitFlag == true {
		return
	} else {
		s.ConnMrg.ClearConn()
		return
	}
}

func (s *Server) AddRouter(msgId uint32, router _interface.Router) {
	s.msgHandler.AddRouter(msgId, router)
	fmt.Println("添加路由成功", router)
}

// GetConnMgr 得到链接管理
func (s *Server) GetConnMgr() _interface.ConnectionManager {
	return s.ConnMrg
}

// SetOnConnStart 设置该Server的连接创建时Hook函数
func (s *Server) SetOnConnStart(hookFunc func(_interface.Connection)) {
	s.OnConnStart = hookFunc
}

// SetOnConnStop 设置该Server的连接断开时的Hook函数
func (s *Server) SetOnConnStop(hookFunc func(_interface.Connection)) {
	s.OnConnStop = hookFunc
}

// CallOnConnStart 调用连接OnConnStart Hook函数
func (s *Server) CallOnConnStart(conn _interface.Connection) {
	if s.OnConnStart != nil {
		fmt.Println("---> 开启回调....")
		s.OnConnStart(conn)
	}
}

// CallOnConnStop 调用连接OnConnStop Hook函数
func (s *Server) CallOnConnStop(conn _interface.Connection) {
	if s.OnConnStop != nil {
		fmt.Println("---> 结束回调....")
		s.OnConnStop(conn)
	}
}

func (s *Server) Packet() _interface.Packet {
	return s.packet
}
