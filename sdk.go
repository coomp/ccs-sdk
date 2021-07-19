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
	RequestHandlers     RequestHandleFuncs
	ResponseHandlers    ResponseHandleFuncs
	messageQueueService *MessageQueueService
}

func NewCcsSdk(requestHandlers RequestHandleFuncs, responseHandlers ResponseHandleFuncs) (*CcsSdk, error) {

	var fixedRequestHandleFuncs RequestHandleFuncs
	fixedRequestHandleFuncs = append(fixedRequestHandleFuncs, NewDefaultRequestHandler())
	fixedRequestHandleFuncs = append(fixedRequestHandleFuncs, requestHandlers...)

	var fixedResponseHandleFuncs ResponseHandleFuncs
	fixedResponseHandleFuncs = append(fixedResponseHandleFuncs, NewDefaultResponseHandler())
	fixedResponseHandleFuncs = append(fixedResponseHandleFuncs, responseHandlers...)

	// TODO fixme
	mqEndpointStr := GetOrDefault(CCS_MQ_ENDPOINT, "118.195.175.6:9876")
	mqEndpoints := strings.SplitN(mqEndpointStr, ",", -1)
	mqReqTopic := GetOrDefault(CCS_MQ_REQ_TOPIC, "__SVC1_REQ_TOPIC__")
	mqRespTopic := GetOrDefault(CCS_MQ_RESP_TOPIC, "__SVC1_REQ_TOPIC__")

	messageQueueRepo, err := NewRocketMQMessageQueueRepository(mqEndpoints)
	if err != nil {
		log.Printf("Could not init rocketmq repo %v", err)
		return nil, err
	}

	sdk := &CcsSdk{
		// TODO fixme
		CcsEndpoint:         GetOrDefault(CCS_ENDPOINTS, "118.195.175.6:2388"),
		RequestHandlers:     fixedRequestHandleFuncs,
		ResponseHandlers:    fixedResponseHandleFuncs,
		messageQueueService: NewMessageQueueService(messageQueueRepo),
	}

	// watch
	err = sdk.messageQueueService.SubscribeRequest(mqReqTopic, sdk.RequestHandlers)
	if err != nil {
		log.Printf("Could not subscribe CCS_MQ_REQ_TOPIC %v", err)
		return nil, err
	}
	err = sdk.messageQueueService.SubscribeResponse(mqRespTopic, sdk.ResponseHandlers)
	if err != nil {
		log.Printf("Could not subscribe CCS_MQ_RESP_TOPIC %v", err)
		return nil, err
	}
	err = sdk.messageQueueService.Start()
	if err != nil {
		log.Printf("Could not start message consumer %v", err)
		return nil, err
	}

	return sdk, nil
}

func (s *CcsSdk) HandleMessage(msg string, need_resp_referers bool, timeout int) error {
	// TODO optimize me, connection multiplexing
	conn, err := grpc.Dial(s.CcsEndpoint, grpc.WithInsecure())

	if err != nil {
		return err
	}

	defer conn.Close()

	messageClient := NewServiceMessageClient(conn)

	secretKey := GetOrDefault(CCS_SECRET_KEY, "__SECRET_KEY__")

	request := &MessageRequest{
		AppId:            GetOrDefault(CCS_APP_ID, "__APP_ID__"),
		ServiceId:        GetOrDefault(CCS_SERVICE_ID, "__SERVICE_ID__"),
		Payload:          msg,
		Timestamp:        time.Now().Unix(),
		NeedRespReferers: need_resp_referers,
	}

	request.Token = HmacSha256Base64(fmt.Sprintf("%s%d", request.AppId, request.Timestamp), secretKey)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	defer cancel()

	_, err = messageClient.OnMessageRequest(ctx, request)
	return err
}
