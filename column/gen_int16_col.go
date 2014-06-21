// Int16Col
// This file was generated. See generator directory.
package column

import (
	"errors"
	"fmt"
)

type Int16Col struct {
	rows []int16
}

func (c *Int16Col) Len() uint {
	return uint(len(c.rows))
}

func (c *Int16Col) Append(row interface{}) error {
	value, found := row.(int16)
	if !found {
		return errors.New(
			fmt.Sprintf("Wrong type (actual) %T != int16 (expected)", row),
		)
	}
	c.rows = append(c.rows, value)
	return nil
}

func (c *Int16Col) At(index uint) (value interface{}, exists bool) {
	if index > uint(len(c.rows)-1) {
		return nil, false
	}
	return c.rows[index], true
}

func (c *Int16Col) String() string {
	return fmt.Sprintf("%#v", c.rows)
}
