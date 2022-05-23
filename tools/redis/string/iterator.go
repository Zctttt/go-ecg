// @Title iterator
// @Description redis 迭代器设计
// @Author zhaocongtao
// @Update 2022/04/25

package string

// Iterator 定义迭代器结构
type Iterator struct {
	// 数据切片
	data []interface{}
	// 当前索引
	index int
}

// 返回迭代器实例
func NewIterator(data []interface{}) *Iterator {
	return &Iterator{
		data: data,
	}
}

// 判断是否有下一切片是否有数据
func (i *Iterator) HasNext() bool {
	return i.getLen() > 0 && i.exist()
}

// 将当前游标走到下一个切片位置 并返回数据
func (i *Iterator) Next() (ret interface{}) {
	ret = i.data[i.index]
	i.index += 1
	return
}

// 获取迭代器数据长度
func (i *Iterator) getLen() int {
	return len(i.data)
}

// 判断迭代器数据可用
func (i *Iterator) exist() bool {
	return i.data != nil
}
