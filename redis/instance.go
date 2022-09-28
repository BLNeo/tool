package redis

import (
	"errors"
	"fmt"
	cluster "github.com/go-redis/redis"
	"github.com/gomodule/redigo/redis"
	"time"
)

type Instance struct {
	DBNumber    int      `toml:"db_number"` //dbnumber
	Addresses   []string `toml:"addresses"` // 连接地址
	Password    string   `toml:"password"`
	MaxIdle     int      `toml:"max_idle"`
	MaxActive   int      `toml:"max_active"`
	IdleTimeout int      `toml:"idle_timeout"`
	Cluster     int      `toml:"cluster"` // 1为集群
}

func (i *Instance) Engine() (*ComponentRedis, error) {
	if len(i.Addresses) == 0 {
		return nil, errors.New("addresses is empty")
	}
	c := &ComponentRedis{
		RedisConn:    nil,
		RedisCluster: nil,
		clusterFlag:  0,
	}
	if i.Cluster == 0 {
		master, err := i.singleRedis()
		if err != nil {
			return nil, err
		}
		c.RedisConn = master
		fmt.Println("使用redis单节点方式")
	} else {
		redisCluster := i.clusterSetup()
		c.RedisCluster = redisCluster
		c.clusterFlag = 1
		fmt.Println("使用redis集群方式")
	}

	return c, nil
}

func (i *Instance) singleRedis() (*redis.Pool, error) {
	conn := &redis.Pool{
		MaxIdle:     i.MaxIdle,
		MaxActive:   i.MaxActive,
		IdleTimeout: time.Duration(i.IdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			fmt.Println("--->redis host :", i.Addresses, i.DBNumber)
			//c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			c, err := redis.Dial("tcp", i.Addresses[0], redis.DialDatabase(i.DBNumber), redis.DialPassword(i.Password))
			if err != nil {
				fmt.Println("redis connect fail")
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return conn, nil
}

//GetRDCluster 集群redis
func (i *Instance) clusterSetup() *cluster.ClusterClient {
	redisCluster := cluster.NewClusterClient(&cluster.ClusterOptions{
		Addrs:    i.Addresses, //set redis cluster url
		Password: i.Password,  //set password
	})
	go redisKeepAliveCluster(redisCluster)
	return redisCluster
}

//redisKeepAliveCluster REDIS 保连
func redisKeepAliveCluster(client *cluster.ClusterClient) {
	for {
		client.Ping().Result()
		time.Sleep(60 * time.Second)
	}
}
