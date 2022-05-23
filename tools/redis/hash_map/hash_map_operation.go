package hash_map

import (
	"context"
	"go-sv/tools/redis/cmd"
)

// StringOperation string类型操作的定义
type HashMapOperation struct {
	ctx context.Context
}

// NewStringOperation 新建string操作实例 隐藏redis本身
func NewHashMapOperation() *HashMapOperation {
	return &HashMapOperation{ctx: context.Background()}
}

func (h HashMapOperation) HSet(key string, collection Collection) {
	cmd.Redis().HSet(h.ctx, key, collection.GetArgs()).Result()
}
