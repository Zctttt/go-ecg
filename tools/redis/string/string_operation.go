// @Title String
// @Description  专门处理string类型的操作
// @Author zhaocongtao
// @Update 2022/04/24

// 业务操作与代码不应该耦合  应该将错误与结果分开

package string

import (
	"context"
	"go-sv/tools/redis/cmd"
	"time"
)

// StringOperation string类型操作的定义
type StringOperation struct {
	ctx context.Context
}

// NewStringOperation 新建string操作实例 隐藏redis本身
func NewStringOperation() *StringOperation {
	return &StringOperation{ctx: context.Background()}
}

// Set 设置string类型
func (s *StringOperation) Set(key string, value interface{}, attrs ...*OperationAttr) *InterfaceResult {
	exp := OperationAtters(attrs).Find(EXPR).UnWarpOr(time.Second * 0).(time.Duration)
	nx := OperationAtters(attrs).Find(NX).UnWarpOr(nil)
	if nx != nil {
		return NewInterfaceResult(cmd.Redis().SetNX(s.ctx, key, value, exp).Result())
	}
	xx := OperationAtters(attrs).Find(XX).UnWarpOr(nil)
	if xx != nil {
		return NewInterfaceResult(cmd.Redis().SetXX(s.ctx, key, value, exp).Result())
	}
	return NewInterfaceResult(cmd.Redis().Set(s.ctx, key, value, exp).Result())
}

// func (s *StringOperation)SetNX
// Get 获取string类型结果
func (s *StringOperation) Get(key string) *StringResult {
	return NewStringResult(cmd.Redis().Get(s.ctx, key).Result())
}

// MGet 获取多个string类型结果
func (s *StringOperation) MGet(key ...string) *SliceResult {
	return NewSliceResult(cmd.Redis().MGet(s.ctx, key...).Result())
}
