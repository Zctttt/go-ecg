package instance

import (
	"fmt"
	"go-sv/cmd/tcp/_interface"
	"sync"
)

func init() {
	CoonMrg = NewConnManager()
}

var CoonMrg *ConnectionManager

type ConnectionManager struct {
	// 管理的连接信息
	connections map[uint32]_interface.Connection
	// 读写链接的读写锁
	coonLock sync.RWMutex
}

func NewConnManager() *ConnectionManager {
	return &ConnectionManager{
		connections: make(map[uint32]_interface.Connection),
	}
}

// Add 添加链接
func (connMgr *ConnectionManager) Add(conn _interface.Connection) {
	// 保护共享资源Map 加写锁
	connMgr.coonLock.Lock()
	// 将 coon 链接添加到 ConnectionManager
	connMgr.connections[conn.GetConnID()] = conn
	connMgr.coonLock.Unlock()

}

// Len 当前连接长度
func (connMgr *ConnectionManager) Len() int {
	connMgr.coonLock.RLock()
	length := len(connMgr.connections)
	connMgr.coonLock.RUnlock()
	return length
}

// Get 利用 ConnID 获取连接
func (connMgr *ConnectionManager) Get(connID uint32) (_interface.Connection, error) {
	connMgr.coonLock.Lock()
	defer connMgr.coonLock.Unlock()

	if coon, ok := connMgr.connections[connID]; ok {
		return coon, nil
	} else {
		return nil, fmt.Errorf("连接未找到 connID = %d", connID)
	}
}

// Remove 删除连接
func (connMgr *ConnectionManager) Remove(conn _interface.Connection) {
	connMgr.coonLock.Lock()
	delete(connMgr.connections, conn.GetConnID())
	connMgr.coonLock.Unlock()
	defer fmt.Println("删除连接成功 ConnID=", conn.GetConnID(), "当前剩余连接数", connMgr.Len())
}

// ClearConn 清除并停止所有连接
func (connMgr *ConnectionManager) ClearConn() {
	connMgr.coonLock.Lock()

	// 停止并清空所有的连接信息
	for connID, conn := range connMgr.connections {
		conn.Stop()
		delete(connMgr.connections, connID)
	}
	connMgr.coonLock.Unlock()
	fmt.Println("所有链接清空 connNum = ", connMgr.Len())
}

//ClearOneConn  利用ConnID获取一个链接 并且删除
func (connMgr *ConnectionManager) ClearOneConn(connID uint32) {
	connMgr.coonLock.Lock()
	defer connMgr.coonLock.Unlock()

	connections := connMgr.connections
	if conn, ok := connections[connID]; ok {
		//停止
		conn.Stop()
		//删除
		delete(connections, connID)
		fmt.Println("Clear Connections ID:  ", connID, "succeed")
		return
	}

	fmt.Println("Clear Connections ID:  ", connID, "err")
	return
}
