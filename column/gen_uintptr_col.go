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

func (c *UintptrCol) Len() int {
	return len(c.rows)
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
