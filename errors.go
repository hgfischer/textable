package textable

import (
	"errors"
	"fmt"
)

const (
	_ERR_INCOR_NUM_VALS string = "Incorrect number of values (actual) %d != %d (expected)"
)

func getError(msg string, values ...interface{}) error {
	return errors.New(fmt.Sprintf(msg, values...))
}
