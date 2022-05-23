// @Title Operation
// @Description  专门处理set操作的附加条件
// @Author zhaocongtao
// @Update 2022/04/24

package string

import (
	"fmt"
	"time"
)

const (
	// 过期时间
	EXPR = "expr"
	// 条件 没有才写入
	NX = "nx"
	// 条件 有才写入
	XX = "xx"
)

type empty struct{}

//属性结构 相对于redis的附加条件
type OperationAttr struct {
	Name  string
	Value interface{}
}

// 表达式切片
type OperationAtters []*OperationAttr

// 加入过期时间
func WithExpire(t time.Duration) *OperationAttr {
	return &OperationAttr{
		Name:  EXPR,
		Value: t,
	}
}

// 条件 不存在才会写入
func WithNX() *OperationAttr {
	return &OperationAttr{
		Name:  NX,
		Value: empty{},
	}
}

// 查找附加条件的表达式
func (attrs OperationAtters) Find(name string) *InterfaceResult {
	for _, attr := range attrs {
		if attr.Name == name {
			return NewInterfaceResult(attr.Value, nil)
		}
	}
	return NewInterfaceResult(nil, fmt.Errorf("operation err : %s", name))
}
