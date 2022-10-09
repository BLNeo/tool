package config

import (
	"github.com/BLNeo/tool/mongo"
	"github.com/BLNeo/tool/mysql"
	"github.com/BLNeo/tool/redis"
	"github.com/BurntSushi/toml"
)

var (
	Config *Conf
)

type Conf struct {
	Mongo mongo.Instance
	Redis redis.Instance
	Mysql mysql.Instance
}

func Init() error {
	config := &Conf{}
	_, err := toml.DecodeFile("./config/conf.toml", config)
	if err != nil {
		return err
	}
	Config = config
	return nil
}
