package rocketmq_impl

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/coomp/ccs-sdk/handle"
	"github.com/coomp/ccs-sdk/message"
)

type RocketMQMessageQueueRepository struct {
	nameServers primitive.NamesrvAddr
	consumer    rocketmq.PushConsumer
	respTopic   string
}

func NewRocketMQMessageQueueRepository(nameServers []string, respTopic string) (*RocketMQMessageQueueRepository, error) {
	log.Printf("NewRocketMQMessageQueueRepository: ns = %v, topic = %v", nameServers, respTopic)

	// mqEndpointStr := env.GetOrDefault(constant.MQ_ENDPOINT, "192.168.147.129:9876")
	// mqEndpoints := strings.SplitN(mqEndpointStr, ",", -1)
	// mqTopic := env.GetOrDefault(constant.MQ_RESP_TOPIC, "coomp")

	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNsResolver(primitive.NewPassthroughResolver(nameServers)),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromFirstOffset),
		consumer.WithConsumerOrder(true),
	)

	return &RocketMQMessageQueueRepository{
		nameServers: nameServers,
		consumer:    c,
		respTopic:   respTopic,
	}, nil
}

func (r RocketMQMessageQueueRepository) Subscribe(topic string, funcs handle.HandleFuncs) {
	err := r.consumer.Subscribe(r.respTopic, consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Printf("On Subscribe Callback: %v \n", msgs[i])
			// msg to context
			var resp message.MessageContextImpl
			err := json.Unmarshal(msgs[i].Body, &resp)
			if err != nil {
				log.Println(err)
			} else {
				for _, handle := range funcs {
					handle(resp)
				}
			}
		}

		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	// Note: start after subscribe
	err = r.consumer.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

}
