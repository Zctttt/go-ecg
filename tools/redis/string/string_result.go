// @Title StringResult
// @Description  专门处理string类型的结果
// @Author zhaocongtao
// @Update 2022/04/24

package string

// StringResult 定义string结果结构体
type StringResult struct {
	Result string
	Err    error
}

// NewStringResult 返回结果结构体
func NewStringResult(result string, err error) *StringResult {
	return &StringResult{
		Result: result,
		Err:    err,
	}
}

// getResult 获取string结果
func (s *StringResult) getResult() string {
	return s.Result
}

// getError 获取错误信息
func (s *StringResult) getError() error {
	return s.Err
}

// UnWarp 解string的结果
func (s *StringResult) UnWarp() string {
	if s.getError() != nil {
		panic(s.getError())
	}
	return s.getResult()
}

// UnWarpOr 根据key解string的结果 如果没有返回当前string
func (s *StringResult) UnWarpOr(key string) string {
	if s.getError() != nil {
		return key
	}
	return s.getResult()
}

// UnWarpOrElse 根据key解string的结果 如果没有执行一个func()
func (s *StringResult) UnWarpOrElse(f func() string) string {
	if s.getError() != nil {
		return f()
	}
	return s.getResult()
}
