package main

import (
	"fmt"
	"log"

	ccssdk "github.com/coomp/ccs-sdk"
)

func NewStdoutRequestLoggingHandle() ccssdk.RequestHandleFunc {
	return func(c ccssdk.RequestMessageContext) error {
		fmt.Printf("DefaultHandle got message at %v\n", c.GetTimestamp())
		return nil
	}
}

func NewStdoutResponseLoggingHandle() ccssdk.ResponseHandleFunc {
	return func(c ccssdk.ResponseMessageContext) error {
		fmt.Printf("DefaultHandle got message at %v\n", c.GetTimestamp())
		return nil
	}
}

func main() {
	requestHandlers := append(ccssdk.NewEmptyRequestHandlers(), NewStdoutRequestLoggingHandle())
	responseHandlers := append(ccssdk.NewEmptyResponseHandlers(), NewStdoutResponseLoggingHandle())

	sdk, err := ccssdk.NewCcsSdk(requestHandlers, responseHandlers)
	if err != nil {
		log.Fatalf("Could not initialize sdk %v", err)
	}

	sdk.HandleMessage("Simple text message", true, 3)

	select {}
}
