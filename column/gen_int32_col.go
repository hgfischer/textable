// Int32Col
// This file was generated. See generator directory.
package column

import (
	"errors"
	"fmt"
)

type Int32Col struct {
	rows []int32
}

func (c *Int32Col) Len() uint {
	return uint(len(c.rows))
}

func (c *Int32Col) Append(row interface{}) error {
	value, found := row.(int32)
	if !found {
		return errors.New(
			fmt.Sprintf("Wrong type (actual) %T != int32 (expected)", row),
		)
	}
	c.rows = append(c.rows, value)
	return nil
}

func (c *Int32Col) At(index uint) (value interface{}, exists bool) {
	if index > uint(len(c.rows)-1) {
		return nil, false
	}
	return c.rows[index], true
}
