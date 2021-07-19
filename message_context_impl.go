package ccssdk

type (
	RequestMessageContextImpl struct {
		AppId            string
		ServiceId        string
		Token            string
		Timestamp        int64
		Payload          string
		Referers         []*interface{}
		NeedRespReferers bool
	}

	ResponseMessageContextImpl struct {
		AppId         string
		ServiceId     string
		RespServiceId string
		Timestamp     int64
		Payload       string
		Referers      []*interface{}
		Code          int32
	}
)

//
// Request
//

func (x RequestMessageContextImpl) GetAppId() string {
	return x.AppId
}

func (x RequestMessageContextImpl) GetServiceId() string {
	return x.ServiceId
}

func (x RequestMessageContextImpl) GetToken() string {
	return x.Token
}

func (x RequestMessageContextImpl) GetTimestamp() int64 {
	return x.Timestamp
}

func (x RequestMessageContextImpl) GetPayload() string {
	return x.Payload
}

func (x RequestMessageContextImpl) GetReferers() []*interface{} {
	return x.Referers
}

func (x RequestMessageContextImpl) GetNeedRespReferers() bool {
	return x.NeedRespReferers
}

//
// Response
//

func (x ResponseMessageContextImpl) GetAppId() string {
	return x.AppId
}

func (x ResponseMessageContextImpl) GetServiceId() string {
	return x.ServiceId
}

func (x ResponseMessageContextImpl) GetRespServiceId() string {
	return x.RespServiceId
}

func (x ResponseMessageContextImpl) GetTimestamp() int64 {
	return x.Timestamp
}

func (x ResponseMessageContextImpl) GetPayload() string {
	return x.Payload
}

func (x ResponseMessageContextImpl) GetReferers() []*interface{} {
	return x.Referers
}

func (x ResponseMessageContextImpl) GetCode() int32 {
	return x.Code
}
