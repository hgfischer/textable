// IntCol
// This file was generated. See generator directory.
package column

import (
	"errors"
	"fmt"
)

type IntCol struct {
	rows []int
}

func (c *IntCol) Len() uint {
	return uint(len(c.rows))
}

func (c *IntCol) Append(row interface{}) error {
	value, found := row.(int)
	if !found {
		return errors.New(
			fmt.Sprintf("Wrong type (actual) %T != int (expected)", row),
		)
	}
	c.rows = append(c.rows, value)
	return nil
}

func (c *IntCol) At(index uint) (value interface{}, exists bool) {
	if index > uint(len(c.rows)-1) {
		return nil, false
	}
	return c.rows[index], true
}
