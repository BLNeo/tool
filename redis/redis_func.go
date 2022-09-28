package redis

import (
	cluster "github.com/go-redis/redis"
	"github.com/gomodule/redigo/redis"
	"sync"
	"time"
)

type ComponentRedis struct {
	RedisConn    *redis.Pool
	RedisCluster *cluster.ClusterClient
	clusterFlag  int
}

var (
	once sync.Once
	cli  *ComponentRedis
)

func GetCli() *ComponentRedis {
	return cli
}

// InitEngine 初始化一个ENGIN
func InitEngine(instance *Instance) (*ComponentRedis, error) {
	engine, err := instance.Engine()
	if err != nil {
		return nil, err
	}
	once.Do(func() {
		cli = engine
	})
	return engine, nil
}

//字符串设置
func (r *ComponentRedis) Set(key, value string) error {
	if r.clusterFlag == 1 {
		err := r.RedisCluster.Set(key, value, 0).Err()
		if err != nil {
			return err
		}
		return nil
	}
	conn := cli.RedisConn.Get()
	defer conn.Close()
	_, err := conn.Do("set", key, value)
	return err
}

//字符串获取
func (r *ComponentRedis) Get(key string) (string, error) {
	if r.clusterFlag == 1 {
		res, err := r.RedisCluster.Get(key).Result()
		if err != nil {
			return "", err
		}
		return res, nil
	} else {
		conn := r.RedisConn.Get()
		defer conn.Close()
		res, err := conn.Do("get", key)
		if err != nil {
			return "", err
		}
		val, err := redis.String(res, err)
		return val, nil
	}
}

//字符串设置
func (r *ComponentRedis) SetAndTimeOut(key, value string, timeout int) error {
	if r.clusterFlag == 1 {
		err := r.RedisCluster.Set(key, value, time.Duration(timeout)*time.Second).Err()
		if err != nil {
			return err
		}
	} else {
		conn := r.RedisConn.Get()
		defer conn.Close()

		_, err := conn.Do("SET", key, value)
		if err != nil {
			return err
		}
		if timeout != 0 {
			_, err = conn.Do("EXPIRE", key, timeout)
			if err != nil {
				return err
			}
		}
		return nil
	}
	return nil
}

//集合获取指定数量的值
func (r *ComponentRedis) SetGetHalf(key, num string) ([]string, error) {
	if r.clusterFlag == 1 {
		res, err := r.RedisCluster.Do("spop", key, num).Result()
		if err != nil {
			return []string{}, err
		}
		return res.([]string), err
	} else {
		conn := r.RedisConn.Get()
		defer conn.Close()

		list, err := redis.Values(conn.Do("spop", key, num))

		if err != nil {
			return []string{}, err
		}
		data := make([]string, 0)
		for _, v := range list {
			data = append(data, string(v.([]byte)))
		}
		return data, nil
	}
}

//集合 添加
func (r *ComponentRedis) SetAddOne(key string, value string) (int64, error) {
	if r.clusterFlag == 1 {
		res, err := r.RedisCluster.SAdd(key, value).Result()
		if err != nil {
			return 0, err
		}

		return res, nil
	} else {
		conn := r.RedisConn.Get()
		defer conn.Close()
		res, err := redis.Int64(conn.Do("SADD", key, value))
		if err != nil {
			return 0, err
		}
		return res, nil
	}
}

// hmget
func (r *ComponentRedis) HMGet(key string, values []string) ([]string, error) {
	if r.clusterFlag == 1 {
		reply, err := r.RedisCluster.HMGet("hmget", values...).Result()
		if err != nil {
			return []string{}, err
		}
		res, err := redis.Strings(reply, err)
		if err != nil {
			return []string{}, err
		}
		return res, nil
	} else {
		conn := r.RedisConn.Get()
		defer conn.Close()

		args := redis.Args{}
		reply, err := conn.Do("hmget", args.Add(key).AddFlat(values)...)
		if err != nil {
			return []string{}, err
		}
		res, err := redis.Strings(reply, err)
		if err != nil {
			return []string{}, err
		}
		return res, nil
	}
}

// hmSet
func (r *ComponentRedis) HMSet(key string, m map[string]string) error {
	if r.clusterFlag == 1 {
		newM := make(map[string]interface{})
		for k, v := range m {
			newM[k] = v
		}
		reply, err := r.RedisCluster.HMSet(key, newM).Result()
		if err != nil {
			return err
		}
		_, err = redis.String(reply, err)
		if err != nil {
			return err
		}
		return nil
	} else {
		conn := r.RedisConn.Get()
		defer conn.Close()
		args := redis.Args{}
		reply, err := conn.Do("hmSet", args.Add(key).AddFlat(m)...)
		if err != nil {
			return err
		}
		_, err = redis.String(reply, err)
		if err != nil {
			return err
		}
		return nil
	}
}
