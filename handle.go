package ccssdk

import (
	"fmt"
)

type (
	HandleFunc  func(MessageContext) error
	HandleFuncs []HandleFunc
)

func NewEmpty() HandleFuncs {
	var handles HandleFuncs
	return handles
}

func NewDefaultHandle() HandleFunc {
	return func(c MessageContext) error {
		fmt.Printf("DefaultHandle got message at %v\n", c.GetTimestamp())
		return nil
	}
}
