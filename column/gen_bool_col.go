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

func (c *BoolCol) Len() int {
	return len(c.rows)
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
