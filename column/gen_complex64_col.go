// Complex64Col
// This file was generated. See generator directory.
package column

import (
	"errors"
	"fmt"
)

type Complex64Col struct {
	rows []complex64
}

func (c *Complex64Col) Len() uint {
	return uint(len(c.rows))
}

func (c *Complex64Col) Append(row interface{}) error {
	value, found := row.(complex64)
	if !found {
		return errors.New(
			fmt.Sprintf("Wrong type (actual) %T != complex64 (expected)", row),
		)
	}
	c.rows = append(c.rows, value)
	return nil
}

func (c *Complex64Col) At(index uint) (value interface{}, exists bool) {
	if index > uint(len(c.rows)-1) {
		return nil, false
	}
	return c.rows[index], true
}

func (c *Complex64Col) String() string {
	return fmt.Sprintf("%#v", c.rows)
}
