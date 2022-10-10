package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// InitEngine 通过配置结构体Instance 实例化出一个连接池mongoClient
func InitEngine(instance *Instance) (*mongo.Client, error) {
	var err error
	URI, err := instance.String()
	fmt.Println(URI)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	clientOptions := options.Client().ApplyURI(URI)
	clientOptions.SetMaxPoolSize(100) //连接池模式
	clientOptions.SetSocketTimeout(60 * time.Second)
	mgoCli, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return mgoCli, nil
}
