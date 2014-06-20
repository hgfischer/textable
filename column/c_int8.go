// Int8Col
// This file was generated. See generator directory.
package column

import (
	"errors"
	"fmt"
)

type Int8Col struct {
	rows []int8
}

func (c *Int8Col) Len() int {
	return len(c.rows)
}

func (c *Int8Col) Append(row interface{}) error {
	value, found := row.(int8)
	if !found {
		return errors.New(
			fmt.Sprintf("Wrong type (actual) %T != int8 (expected)", row),
		)
	}
	c.rows = append(c.rows, value)
	return nil
}
