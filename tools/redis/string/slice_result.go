// @Title SliceResult
// @Description  专门处理slice类型的结果
// @Author zhaocongtao
// @Update 2022/04/24

package string

// SliceResult 定义slice结果结构体
type SliceResult struct {
	Result []interface{}
	Err    error
}

// NewSliceResult 返回结果结构体
func NewSliceResult(result []interface{}, err error) *SliceResult {
	return &SliceResult{
		Result: result,
		Err:    err,
	}
}

// getResult 获取slice结果
func (s *SliceResult) getResult() []interface{} {
	return s.Result
}

// getError 获取错误信息
func (s *SliceResult) getError() error {
	return s.Err
}

// UnWarp 解slice的结果
func (s *SliceResult) UnWarp() []interface{} {
	if s.getError() != nil {
		panic(s.getError())
	}
	return s.getResult()
}

// UnWarpOr 根据keys解slice的结果 如果没有返回当前string
func (s *SliceResult) UnWarpOr(v []interface{}) []interface{} {
	if s.getError() != nil {
		return v
	}
	return s.getResult()
}

// 获取当前切片的迭代器
func (s *SliceResult) Iter() *Iterator {
	return NewIterator(s.Result)
}
