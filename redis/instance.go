package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"sync"
)

type Instance struct {
	DBNumber    int      `toml:"db_number"` //
	Addresses   []string `toml:"addresses"` // 连接地址
	Password    string   `toml:"password"`
	MaxIdle     int      `toml:"max_idle"`
	MaxActive   int      `toml:"max_active"`
	IdleTimeout int      `toml:"idle_timeout"`
}

var (
	once sync.Once
	rdb  *redis.Client
)

func GetRdb() *redis.Client {
	return rdb
}

// InitClient 初始化一个ENGIN
func InitClient(i *Instance) error {
	if len(i.Addresses) == 0 {
		return errors.New("addresses is empty")
	}

	client := redis.NewClient(&redis.Options{
		Addr:           i.Addresses[0],
		Password:       "", // 没有密码，默认值
		DB:             0,  // 默认DB 0
		MaxActiveConns: i.MaxActive,
		MaxIdleConns:   i.MaxIdle,
	})

	err := client.Ping(context.Background()).Err()
	if err != nil {
		return err
	}
	once.Do(func() {
		rdb = client
	})

	fmt.Println("使用redis单节点方式")

	return nil
}
