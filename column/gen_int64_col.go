// Int64Col
// This file was generated. See generator directory.
package column

import (
	"errors"
	"fmt"
)

type Int64Col struct {
	rows []int64
}

func (c *Int64Col) Len() uint {
	return uint(len(c.rows))
}

func (c *Int64Col) Append(row interface{}) error {
	value, found := row.(int64)
	if !found {
		return errors.New(
			fmt.Sprintf("Wrong type (actual) %T != int64 (expected)", row),
		)
	}
	c.rows = append(c.rows, value)
	return nil
}

func (c *Int64Col) At(index uint) (value interface{}, exists bool) {
	if index > uint(len(c.rows)-1) {
		return nil, false
	}
	return c.rows[index], true
}
