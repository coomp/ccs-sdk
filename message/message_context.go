package message

type MessageContext interface {
	GetAppId() string
	GetServiceId() string
	GetRespServiceId() string
	GetTimestamp() int64
	GetPayload() string
	GetReferers() []*interface{}
	GetCode() int32
}
