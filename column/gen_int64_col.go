// Int64Col
// This file was generated. See generator directory.
package column

import (
	"errors"
	"fmt"
)

type Int64Col struct {
	rows []int64
}

func (c *Int64Col) Len() int {
	return len(c.rows)
}

func (c *Int64Col) Append(row interface{}) error {
	value, found := row.(int64)
	if !found {
		return errors.New(
			fmt.Sprintf("Wrong type (actual) %T != int64 (expected)", row),
		)
	}
	c.rows = append(c.rows, value)
	return nil
}
