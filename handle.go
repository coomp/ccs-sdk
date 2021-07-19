package ccssdk

import (
	"fmt"
)

type (
	RequestHandleFunc  func(RequestMessageContext) error
	RequestHandleFuncs []RequestHandleFunc
)

func NewEmptyRequestHandlers() RequestHandleFuncs {
	var handles RequestHandleFuncs
	return handles
}

func NewDefaultRequestHandler() RequestHandleFunc {
	return func(c RequestMessageContext) error {
		fmt.Printf("DefaultRequestHandler got message at %v\n", c.GetTimestamp())
		return nil
	}
}

type (
	ResponseHandleFunc  func(ResponseMessageContext) error
	ResponseHandleFuncs []ResponseHandleFunc
)

func NewEmptyResponseHandlers() ResponseHandleFuncs {
	var handles ResponseHandleFuncs
	return handles
}

func NewDefaultResponseHandler() ResponseHandleFunc {
	return func(c ResponseMessageContext) error {
		fmt.Printf("DefaultResponseHandler got message at %v\n", c.GetTimestamp())
		return nil
	}
}
