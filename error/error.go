package error

import (
	"errors"
	"fmt"
)

var (
	BaseError  = errors.New("base error")
	SuperError = fmt.Errorf("%w: super", BaseError)
)

func WrapErrorFromBase(message string) error {
	return fmt.Errorf("%w: %v", BaseError, message)
}

func WrapErrorFromSuper(message string) error {
	return fmt.Errorf("%w: %v", SuperError, message)
}
