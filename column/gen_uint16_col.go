// Uint16Col
// This file was generated. See generator directory.
package column

import (
	"errors"
	"fmt"
)

type Uint16Col struct {
	rows []uint16
}

func (c *Uint16Col) Len() int {
	return len(c.rows)
}

func (c *Uint16Col) Append(row interface{}) error {
	value, found := row.(uint16)
	if !found {
		return errors.New(
			fmt.Sprintf("Wrong type (actual) %T != uint16 (expected)", row),
		)
	}
	c.rows = append(c.rows, value)
	return nil
}
