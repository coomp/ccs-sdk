package main

import (
	"fmt"

	"github.com/coomp/ccs-sdk"
	"github.com/coomp/ccs-sdk/handle"
	"github.com/coomp/ccs-sdk/message"
)

func NewStdoutLoggingHandle() handle.HandleFunc {
	return func(c message.MessageContext) error {
		fmt.Printf("DefaultHandle got message at %v\n", c.GetTimestamp())
		return nil
	}
}

func main() {
	handles := handle.NewEmpty()
	handles = append(handles, NewStdoutLoggingHandle())
	ccs.InitCcsSdk(handles)
}
