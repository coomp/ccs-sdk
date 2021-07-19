package handle

import (
	"fmt"

	"github.com/coomp/ccs-sdk/message"
)

type (
	HandleFunc  func(message.MessageContext) error
	HandleFuncs []HandleFunc
)

func NewEmpty() HandleFuncs {
	var handles HandleFuncs
	return handles
}

func NewDefaultHandle() HandleFunc {
	return func(c message.MessageContext) error {
		fmt.Printf("DefaultHandle got message at %v\n", c.GetTimestamp())
		return nil
	}
}
