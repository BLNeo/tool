package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"os"
	"os/signal"
	"testing"
	"time"
)

// InitKafka 初始化kafka队列
func InitKafka() *Kafka {

	ins := &Instance{
		Address: []string{"127.0.0.1:9092"},
		Topics:  []string{"test1", "test2"},
	}

	// 这里是producer
	producerConf := sarama.NewConfig()
	producerConf.Producer.Return.Successes = true
	producerConf.Producer.Partitioner = sarama.NewHashPartitioner
	producerConf.Producer.Timeout = 5 * time.Second
	producerConf.Producer.MaxMessageBytes = 1024000 //最大发送消息体大小 1024kb 1M
	ins.Producer = producerConf

	// 这里是consumer
	ComsumerConf := cluster.NewConfig()
	ComsumerConf.Consumer.Return.Errors = true
	ComsumerConf.Consumer.Offsets.Initial = sarama.OffsetNewest // 从记录中的偏移量消费
	ComsumerConf.Group.Return.Notifications = true
	ins.Consumer = ComsumerConf

	// 进行测试
	return NewKafka(ins)
}

// 测试生产者
func TestProducer(t *testing.T) {
	k := InitKafka()
	producer, err := k.Producer()
	if err != nil {
		t.Fatal(err)
	}
	go handlerProducer(producer, k.ins.Topics[0])
	go handlerProducer(producer, k.ins.Topics[1])
	select {}
}

// 测试消费者
func TestConsumer(t *testing.T) {
	k := InitKafka()
	// 此处也可以针对topic进行分开  用goroutine消费对应的
	consumer, err := k.Consumer(k.ins.Topics, "group1")
	if err != nil {
		t.Fatal(err)
	}
	handleConsumer(consumer)

	select {}
}

//handlerProducer 处理生产者函数
func handlerProducer(p sarama.AsyncProducer, topic string) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

		for {
			select {
			case err := <-p.Errors():
				fmt.Println(err)
			case succ := <-p.Successes():
				fmt.Println(fmt.Sprintf("发送成功 ： topic: %s partition: %d offset : %d value: %s", succ.Topic, succ.Partition, succ.Offset, succ.Value))
			}
		}
	}()

	// 推送内容到对应具体的topic
	for {
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder("hello world"),
		}
		p.Input() <- msg
		time.Sleep(200 * time.Millisecond)
	}
}

//handlerProducer 处理生产者函数
func handleConsumer(c *cluster.Consumer) {
	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	// consume errors
	go func() {
		for err := range c.Errors() {
			fmt.Println(fmt.Sprintf("Consumer发生错误：%v", err))
		}
	}()

	// consume notifications
	go func() {
		for info := range c.Notifications() {
			fmt.Println("kafka 重平衡 : " + fmt.Sprintf("%+v", info))
		}
	}()

	for {
		select {
		case msg, ok := <-c.Messages():
			if ok {
				fmt.Println(fmt.Sprintf("接收消息 ： topic: %s partition: %d offset : %d ", msg.Topic, msg.Partition, msg.Offset))
				// 业务处理。。。
				c.MarkOffset(msg, "") // mark message as processed
			}
		case <-signals:
			return
		}
	}
}
