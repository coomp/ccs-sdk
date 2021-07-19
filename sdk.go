package ccs

import (
	"log"
	"strings"

	"github.com/coomp/ccs-sdk/constant"
	"github.com/coomp/ccs-sdk/env"
	"github.com/coomp/ccs-sdk/handle"
	"github.com/coomp/ccs-sdk/repositories/rocketmq_impl"
	"github.com/coomp/ccs-sdk/services"
)

type CcsSdk struct {
	HandleFuncs         handle.HandleFuncs
	messageQueueService *services.MessageQueueService
}

func InitCcsSdk(handleFuncs handle.HandleFuncs) (*CcsSdk, error) {

	var fixedHandleFuncs handle.HandleFuncs
	fixedHandleFuncs = append(fixedHandleFuncs, handle.NewDefaultHandle())
	fixedHandleFuncs = append(fixedHandleFuncs, handleFuncs...)

	// mq
	mqEndpointStr := env.GetOrDefault(constant.MQ_ENDPOINT, "192.168.147.129:9876")
	mqEndpoints := strings.SplitN(mqEndpointStr, ",", -1)
	mqRespTopic := env.GetOrDefault(constant.MQ_RESP_TOPIC, "svc1-resp-topic")

	messageQueueRepo, err := rocketmq_impl.NewRocketMQMessageQueueRepository(mqEndpoints, mqRespTopic)
	if err != nil {
		log.Fatal(err)
	}

	sdk := &CcsSdk{
		HandleFuncs:         fixedHandleFuncs,
		messageQueueService: services.NewMessageQueueService(messageQueueRepo),
	}

	// watch
	sdk.messageQueueService.Subscribe(mqRespTopic, sdk.HandleFuncs)

	return sdk, nil
}
