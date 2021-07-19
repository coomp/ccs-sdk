package main

import (
	"fmt"

	ccssdk "github.com/coomp/ccs-sdk"
)

func NewStdoutLoggingHandle() ccssdk.HandleFunc {
	return func(c ccssdk.MessageContext) error {
		fmt.Printf("DefaultHandle got message at %v\n", c.GetTimestamp())
		return nil
	}
}

func main() {
	handles := ccssdk.NewEmpty()
	handles = append(handles, NewStdoutLoggingHandle())
	ccssdk.InitCcsSdk(handles)
}
