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

func (c *UintCol) Len() int {
	return len(c.rows)
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
