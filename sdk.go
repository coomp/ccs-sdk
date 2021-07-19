package ccssdk

import (
	context "context"
	"fmt"
	"log"
	"strings"
	"time"

	grpc "google.golang.org/grpc"
)

type CcsSdk struct {
	CcsEndpoint         string
	HandleFuncs         HandleFuncs
	messageQueueService *MessageQueueService
}

func InitCcsSdk(ccs_endpoint string, handleFuncs HandleFuncs) (*CcsSdk, error) {

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
		CcsEndpoint:         ccs_endpoint,
		HandleFuncs:         fixedHandleFuncs,
		messageQueueService: NewMessageQueueService(messageQueueRepo),
	}

	// watch
	sdk.messageQueueService.Subscribe(mqRespTopic, sdk.HandleFuncs)

	return sdk, nil
}

func (s *CcsSdk) Send(msg string, timeout int) error {
	// TODO optimize me, connection multiplexing
	conn, err := grpc.Dial(s.CcsEndpoint, grpc.WithInsecure())

	if err != nil {
		return err
	}

	defer conn.Close()

	messageClient := NewServiceMessageClient(conn)

	// TODO from env
	secretKey := "__SECRET_KEY__"

	// TODO fill AppId, ServiceId etc.
	request := &MessageRequest{
		AppId:     "__APP_ID__",
		ServiceId: "__SERVICE_ID__",
		Payload:   msg,
		Timestamp: time.Now().Unix(),
	}

	request.Token = HmacSha256Base64(fmt.Sprintf("%s%d", request.AppId, request.Timestamp), secretKey)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	defer cancel()

	_, err = messageClient.OnMessageRequest(ctx, request)
	return err
}
