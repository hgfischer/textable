// Uint64Col
// This file was generated. See generator directory.
package column

import (
	"errors"
	"fmt"
)

type Uint64Col struct {
	rows []uint64
}

func (c *Uint64Col) Len() int {
	return len(c.rows)
}

func (c *Uint64Col) Append(row interface{}) error {
	value, found := row.(uint64)
	if !found {
		return errors.New(
			fmt.Sprintf("Wrong type (actual) %T != uint64 (expected)", row),
		)
	}
	c.rows = append(c.rows, value)
	return nil
}
