// Uint8Col
// This file was generated. See generator directory.
package column

import (
	"errors"
	"fmt"
)

type Uint8Col struct {
	rows []uint8
}

func (c *Uint8Col) Len() int {
	return len(c.rows)
}

func (c *Uint8Col) Append(row interface{}) error {
	value, found := row.(uint8)
	if !found {
		return errors.New(
			fmt.Sprintf("Wrong type (actual) %T != uint8 (expected)", row),
		)
	}
	c.rows = append(c.rows, value)
	return nil
}
