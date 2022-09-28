package redis

import (
	"fmt"
	"testing"
)

//[redis]
//Addresses = "127.0.0.1:6379"
//Password = ""
//MaxIdle = 30
//MaxActive = 30
//IdleTimeout = 200
//Cluster = 0
//DBNumber = 7

func TestRedis(t *testing.T) {
	ins := &Instance{
		DBNumber:    0,
		Addresses:   []string{"127.0.0.1:6379"},
		Password:    "",
		MaxIdle:     30,
		MaxActive:   30,
		IdleTimeout: 200,
		Cluster:     0,
	}
	redisEngine, err := InitEngine(ins)
	if err != nil {
		t.Fatal(err)
	}
	err = redisEngine.Set("test1", "hello world")
	if err != nil {
		t.Fatal(err)
	}
	value, err := redisEngine.Get("test1")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(value)
}
