package ccssdk

type RequestMessageContext interface {
	GetAppId() string
	GetServiceId() string
	GetToken() string
	GetTimestamp() int64
	GetPayload() string
	GetReferers() []*interface{}
	GetNeedRespReferers() bool
}

type ResponseMessageContext interface {
	GetAppId() string
	GetServiceId() string
	GetRespServiceId() string
	GetTimestamp() int64
	GetPayload() string
	GetReferers() []*interface{}
	GetCode() int32
}
