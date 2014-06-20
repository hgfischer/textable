// RuneCol
// This file was generated. See generator directory.
package column

import (
	"errors"
	"fmt"
)

type RuneCol struct {
	rows []rune
}

func (c *RuneCol) Len() int {
	return len(c.rows)
}

func (c *RuneCol) Append(row interface{}) error {
	value, found := row.(rune)
	if !found {
		return errors.New(
			fmt.Sprintf("Wrong type (actual) %T != rune (expected)", row),
		)
	}
	c.rows = append(c.rows, value)
	return nil
}
