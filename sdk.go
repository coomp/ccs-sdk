package ccssdk

import (
	"log"
	"strings"
)

type CcsSdk struct {
	HandleFuncs         HandleFuncs
	messageQueueService *MessageQueueService
}

func InitCcsSdk(handleFuncs HandleFuncs) (*CcsSdk, error) {

	var fixedHandleFuncs HandleFuncs
	fixedHandleFuncs = append(fixedHandleFuncs, NewDefaultHandle())
	fixedHandleFuncs = append(fixedHandleFuncs, handleFuncs...)

	// mq
	mqEndpointStr := GetOrDefault(MQ_ENDPOINT, "192.168.147.129:9876")
	mqEndpoints := strings.SplitN(mqEndpointStr, ",", -1)
	mqRespTopic := GetOrDefault(MQ_RESP_TOPIC, "svc1-resp-topic")

	messageQueueRepo, err := NewRocketMQMessageQueueRepository(mqEndpoints, mqRespTopic)
	if err != nil {
		log.Fatal(err)
	}

	sdk := &CcsSdk{
		HandleFuncs:         fixedHandleFuncs,
		messageQueueService: NewMessageQueueService(messageQueueRepo),
	}

	// watch
	sdk.messageQueueService.Subscribe(mqRespTopic, sdk.HandleFuncs)

	return sdk, nil
}
