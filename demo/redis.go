// @Title
// @Description
// @Author
// @Update

package main

import (
	"context"
	"fmt"
	"go-sv/tools/redis/cmd"
)

func main() {
	ctx := context.Background()
	ret := cmd.Redis().Get(ctx, "foo")
	v, err := ret.Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}
