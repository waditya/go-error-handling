package customerror

import (
	"errors"
	"fmt"
)

func NewError(errorString string) error {
	err := errors.New(errorString)
	return err
}

func NewErrorVerbose(errorString string, number int64) error {
	err := fmt.Errorf(errorString+" got : %d", number)
	return err
}
