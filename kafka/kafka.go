package kafka

import (
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

type Kafka struct {
	ins *Instance
}

func NewKafka(instance *Instance) *Kafka {
	return &Kafka{ins: instance}
}

//Producer 获取到一个Producer,但是记住要进行P的多个管道进行管理，具体可参照test案例
// 在调用Produce发送内容之前，你需要知道： 要发送到哪个topic中，要发送的是什么数据
// 记得要关掉这个连接在程序释放的时候
func (k *Kafka) Producer() (sarama.AsyncProducer, error) {
	return sarama.NewAsyncProducer(k.ins.Address, k.ins.Producer)
}

// Consumer 这里会包含一个Group概念，什么是Group呢，一般来说说Group就是指的消费者分组
// 如果说有两个业务组，A组、B组同时消耗这份数据，就需要在定义消费者的时候去消耗内容，
// 记得要关掉这个连接在程序释放的时候
func (k *Kafka) Consumer(topics []string, group string) (*cluster.Consumer, error) {
	if group == "" {
		group = k.ins.Group
	}
	return cluster.NewConsumer(k.ins.Address, group, topics, k.ins.Consumer)
}
