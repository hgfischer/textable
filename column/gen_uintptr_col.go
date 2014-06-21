// UintptrCol
// This file was generated. See generator directory.
package column

import (
	"errors"
	"fmt"
)

type UintptrCol struct {
	rows []uintptr
}

func (c *UintptrCol) Len() uint {
	return uint(len(c.rows))
}

func (c *UintptrCol) Append(row interface{}) error {
	value, found := row.(uintptr)
	if !found {
		return errors.New(
			fmt.Sprintf("Wrong type (actual) %T != uintptr (expected)", row),
		)
	}
	c.rows = append(c.rows, value)
	return nil
}

func (c *UintptrCol) At(index uint) (value interface{}, exists bool) {
	if index > uint(len(c.rows)-1) {
		return nil, false
	}
	return c.rows[index], true
}

func (c *UintptrCol) String() string {
	return fmt.Sprintf("%#v", c.rows)
}
