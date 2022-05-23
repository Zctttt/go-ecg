package ecg

import (
	"context"
	"crypto/rand"
	"fmt"
	"go-sv/models"
	redisCmd "go-sv/tools/redis/cmd"
	redisStr "go-sv/tools/redis/string"
	"math/big"
	"time"
)

type York struct {
	ctx context.Context
}

func NewYork() *York {
	return &York{
		ctx: context.Background(),
	}
}
func randUUID() [4]byte {
	var UUID = [4]byte{0, 0, 0, 0}
	for i := 0; i < 4; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(255))
		UUID[i] = byte(n.Int64())
	}
	return UUID
}

// 首次tcp链接分配uuid
func (y *York) Connect() [4]byte {
	Key := randUUID()
	KeySlice := Key[0:4]
	redisStr.NewStringOperation().Set(string(KeySlice), 1, redisStr.WithExpire(time.Second*1000))
	redisCmd.Redis().SAdd(y.ctx, "user", string(KeySlice))
	fmt.Println(Key)
	return Key
}
func (y *York) UserList() []byte {
	user, err := redisCmd.Redis().SMembers(y.ctx, "user").Result()
	if err != nil {
		return nil
	}
	result := redisStr.NewStringOperation().MGet(user...).UnWarp()
	b := make([]byte, 0)
	for k, v := range result {
		if bt, ok := v.(string); ok {
			if bt == "1" {
				b = append(b, user[k]...)
			}

		} else {
			fmt.Print("user_list err data :")
			fmt.Println(bt)
		}
	}
	return b
}

// udp传输并缓存
func (y *York) Set(data []byte) bool {
	h := models.NewHeartData(data)
	redisStr.NewStringOperation().Set(string(append(h.UUID, h.TIME...)), data, redisStr.WithExpire(time.Minute*10), redisStr.WithNX())
	//fmt.Println(redisCmd.Redis().ZAdd(y.ctx, string(h.UID), &redis.Z{
	//	Score:  float64(binary.LittleEndian.Uint32(h.TIME)), //十六进制转float64
	//	Member: h,
	//}).Result())
	return true
}

func (y *York) Get(uuid string, time string) (*models.HeartData, error) {
	result := redisStr.NewStringOperation().Get(uuid + time)
	if result.Err != nil {
		return nil, result.Err
	}
	return models.NewHeartData([]byte(result.Result)), result.Err
}
