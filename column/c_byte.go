// ByteCol
// This file was generated. See generator directory.
package column

import (
	"errors"
	"fmt"
)

type ByteCol struct {
	rows []byte
}

func (c *ByteCol) Len() int {
	return len(c.rows)
}

func (c *ByteCol) Append(row interface{}) error {
	value, found := row.(byte)
	if !found {
		return errors.New(
			fmt.Sprintf("Wrong type (actual) %T != byte (expected)", row),
		)
	}
	c.rows = append(c.rows, value)
	return nil
}
