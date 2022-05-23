package utils

import (
	"encoding/json"
	"go-sv/cmd/tcp/_interface"
	"io/ioutil"
)

type GlobalObj struct {
	// 全局Server对象
	TcpServer _interface.Server
	// 当前主机IP
	Host string
	// 当前主机端口
	Port int
	// 当前服务器名称
	Name string

	// 当前框架版本
	Version string
	// 当前服务器允许的最大连接数
	MaxConn int
	// 当前框架数据包最大值
	MaxPackageSize uint32
	//业务工作Worker池的数量
	WorkerPoolSize uint32
	//业务工作Worker对应负责的任务队列最大任务存储数量
	MaxWorkerTaskLen uint32

	MaxMsgChanLen uint32
	// 配置文件路径
	ConfFilePath string
}

var GlobalObject *GlobalObj

// Reload 加载自定义参数
func (g *GlobalObj) Reload() {
	if data, err := ioutil.ReadFile("conf/develop.json"); err != nil {
		panic(err)
	} else {
		if err := json.Unmarshal(data, &GlobalObject); err != nil {
			panic(err)
		}
	}
}

func init() {
	GlobalObject = &GlobalObj{
		//TcpServer:      nil,
		Host:             "0.0.0.0",
		Port:             9000,
		Name:             "Basic",
		Version:          "0.1",
		MaxConn:          1000,
		MaxPackageSize:   4096,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
		ConfFilePath:     "conf/develop.json",
	}
	// 加载自定义参数
	GlobalObject.Reload()
}
