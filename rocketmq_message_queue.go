package ccssdk

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

type RocketMQMessageQueueRepository struct {
	nameServers primitive.NamesrvAddr
	consumer    rocketmq.PushConsumer
}

func NewRocketMQMessageQueueRepository(nameServers []string) (*RocketMQMessageQueueRepository, error) {
	log.Printf("NewRocketMQMessageQueueRepository: ns = %v", nameServers)

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
	}, nil
}

func (r RocketMQMessageQueueRepository) SubscribeRequest(topic string, funcs RequestHandleFuncs) error {
	err := r.consumer.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Printf("On SubscribeRequest Callback: %v \n", msgs[i])
			// msg to context
			var req RequestMessageContextImpl
			err := json.Unmarshal(msgs[i].Body, &req)
			if err != nil {
				log.Println(err)
			} else {
				for _, handle := range funcs {
					handle(req)
				}
			}
		}

		return consumer.ConsumeSuccess, nil
	})
	return err

}

func (r RocketMQMessageQueueRepository) SubscribeResponse(topic string, funcs ResponseHandleFuncs) error {
	err := r.consumer.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Printf("On SubscribeResponse Callback: %v \n", msgs[i])
			// msg to context
			var resp ResponseMessageContextImpl
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

	return err

}

func (r RocketMQMessageQueueRepository) Start() error {
	// Note: start after subscribe
	return r.consumer.Start()
}
