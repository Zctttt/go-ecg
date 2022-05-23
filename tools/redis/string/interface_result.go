// @Title interface
// @Description redis 迭代器设计
// @Author zhaocongtao
// @Update 2022/04/25

package string

type InterfaceResult struct {
	Result interface{}
	Err    error
}

func NewInterfaceResult(result interface{}, err error) *InterfaceResult {
	return &InterfaceResult{
		Result: result,
		Err:    err,
	}
}

// getResult 获取string结果
func (s *InterfaceResult) getResult() interface{} {
	return s.Result
}

// getError 获取错误信息
func (s *InterfaceResult) getError() error {
	return s.Err
}

// UnWarp 解string的结果
func (s *InterfaceResult) UnWarp() interface{} {
	if s.getError() != nil {
		panic(s.getError())
	}
	return s.getResult()
}

// UnWarpOr 根据key解string的结果 如果没有返回当前string
func (s *InterfaceResult) UnWarpOr(v interface{}) interface{} {
	if s.getError() != nil {
		return v
	}
	return s.getResult()
}
