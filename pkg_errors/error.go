package error

import (
	"fmt"

	"github.com/pkg/errors"
)

var (
	BaseError     = errors.New("base error")
	SuperError    = fmt.Errorf("%w: super", BaseError)
	OriginalError = errors.New("original error")
)

func WrapErrorFromBase(message string) error {
	return fmt.Errorf("%w: %v", BaseError, message)
}

func WrapErrorFromSuper(message string) error {
	return fmt.Errorf("%w: %v", SuperError, message)
}

func WrapOriginalErrorFromBase() error {
	return fmt.Errorf("%+v: %w", OriginalError, BaseError)
}

func WrapBaseErrorFromOriginal() error {
	return fmt.Errorf("%+v: %w", BaseError, OriginalError)
}
