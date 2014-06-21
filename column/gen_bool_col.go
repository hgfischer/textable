// BoolCol
// This file was generated. See generator directory.
package column

import (
	"errors"
	"fmt"
)

type BoolCol struct {
	rows []bool
}

func (c *BoolCol) Len() uint {
	return uint(len(c.rows))
}

func (c *BoolCol) Append(row interface{}) error {
	value, found := row.(bool)
	if !found {
		return errors.New(
			fmt.Sprintf("Wrong type (actual) %T != bool (expected)", row),
		)
	}
	c.rows = append(c.rows, value)
	return nil
}

func (c *BoolCol) At(index uint) (value interface{}, exists bool) {
	if index > uint(len(c.rows)-1) {
		return nil, false
	}
	return c.rows[index], true
}

func (c *BoolCol) String() string {
	return fmt.Sprintf("%#v", c.rows)
}
