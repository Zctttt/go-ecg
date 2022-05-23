package instance

import "go-sv/cmd/tcp/_interface"

// BaseRoute 基础路由 不需要全部重新接口 在router嵌入BaseRouter重写方法就可以
type BaseRouter struct {
	Service interface{}
}

func (b *BaseRouter) PreHandle(request _interface.Request) {
}

func (b *BaseRouter) Handle(request _interface.Request) {
}

func (b *BaseRouter) PostHandle(request _interface.Request) {
}
