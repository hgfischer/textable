// Complex128Col
// This file was generated. See generator directory.
package column

import (
	"errors"
	"fmt"
)

type Complex128Col struct {
	rows []complex128
}

func (c *Complex128Col) Len() int {
	return len(c.rows)
}

func (c *Complex128Col) Append(row interface{}) error {
	value, found := row.(complex128)
	if !found {
		return errors.New(
			fmt.Sprintf("Wrong type (actual) %T != complex128 (expected)", row),
		)
	}
	c.rows = append(c.rows, value)
	return nil
}
