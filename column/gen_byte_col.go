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

func (c *ByteCol) Len() uint {
	return uint(len(c.rows))
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

func (c *ByteCol) At(index uint) (value interface{}, exists bool) {
	if index > uint(len(c.rows)-1) {
		return nil, false
	}
	return c.rows[index], true
}

func (c *ByteCol) String() string {
	return fmt.Sprintf("%#v", c.rows)
}
