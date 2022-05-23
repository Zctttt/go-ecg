package _interface

// Router 路由抽象层
type Router interface {
	// PreHandle 处理业务前钩子
	PreHandle(request Request)
	// Handle 处理业务主方法
	Handle(request Request)
	// PostHandle 处理业务后钩子
	PostHandle(request Request)
}
