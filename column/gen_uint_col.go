// UintCol
// This file was generated. See generator directory.
package column

import (
	"errors"
	"fmt"
)

type UintCol struct {
	rows []uint
}

func (c *UintCol) Len() uint {
	return uint(len(c.rows))
}

func (c *UintCol) Append(row interface{}) error {
	value, found := row.(uint)
	if !found {
		return errors.New(
			fmt.Sprintf("Wrong type (actual) %T != uint (expected)", row),
		)
	}
	c.rows = append(c.rows, value)
	return nil
}

func (c *UintCol) At(index uint) (value interface{}, exists bool) {
	if index > uint(len(c.rows)-1) {
		return nil, false
	}
	return c.rows[index], true
}

func (c *UintCol) String() string {
	return fmt.Sprintf("%#v", c.rows)
}
