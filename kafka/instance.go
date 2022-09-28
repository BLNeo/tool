package kafka

import (
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

// 生产者相关设置
// config.Producer.Return.Successes = true
// config.Producer.Partitioner = sarama.NewHashPartitioner
// config.Producer.Timeout = 5 * time.Second
//	config.Producer.MaxMessageBytes = config.Kafka.MaxMessageBytes //最大发送消息体大小设置为2M，单位byte

// 群组相关设置
// config.Group.Return.Notifications = true

// 这里是消费者相关设置
//	config.Consumer.Return.Errors = true
// config.Consumer.Offsets.Initial = sarama.OffsetNewest

type Instance struct {
	Group           string   `toml:"group"`             // 分组
	Address         []string `toml:"address"`           // 连接地址
	Topics          []string `toml:"topics"`            // 所有topic
	MaxMessageBytes int      `toml:"max_message_bytes"` // //最大发送消息体大小设置，单位byte
	Producer        *sarama.Config
	Consumer        *cluster.Config
}
