package test

import (
	"context"
	"fmt"
	"go-sv/tools/redis/cmd"
	"testing"
)

//func TestSet(tv *testing.T) {
//	fmt.Println(string2.NewStringOperation().Set("hello", "world", string2.WithExpire(10*time.Second)))
//}

func TestSet(tv *testing.T) {
	fmt.Println(cmd.Redis().SMembers(context.Background(), "user").Result())
}
