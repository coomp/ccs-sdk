package ccssdk

type MessageContextImpl struct {
	AppId         string
	ServiceId     string
	RespServiceId string
	Timestamp     int64
	Payload       string
	Referers      []*interface{}
	Code          int32
}

func (x MessageContextImpl) GetAppId() string {
	return x.AppId
}

func (x MessageContextImpl) GetServiceId() string {
	return x.ServiceId
}

func (x MessageContextImpl) GetRespServiceId() string {
	return x.RespServiceId
}

func (x MessageContextImpl) GetTimestamp() int64 {
	return x.Timestamp
}

func (x MessageContextImpl) GetPayload() string {
	return x.Payload
}

func (x MessageContextImpl) GetReferers() []*interface{} {
	return x.Referers
}

func (x MessageContextImpl) GetCode() int32 {
	return x.Code
}
