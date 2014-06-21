// Uint32Col
// This file was generated. See generator directory.
package column

import (
	"errors"
	"fmt"
)

type Uint32Col struct {
	rows []uint32
}

func (c *Uint32Col) Len() uint {
	return uint(len(c.rows))
}

func (c *Uint32Col) Append(row interface{}) error {
	value, found := row.(uint32)
	if !found {
		return errors.New(
			fmt.Sprintf("Wrong type (actual) %T != uint32 (expected)", row),
		)
	}
	c.rows = append(c.rows, value)
	return nil
}

func (c *Uint32Col) At(index uint) (value interface{}, exists bool) {
	if index > uint(len(c.rows)-1) {
		return nil, false
	}
	return c.rows[index], true
}
