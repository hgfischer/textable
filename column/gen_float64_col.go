// Float64Col
// This file was generated. See generator directory.
package column

import (
	"errors"
	"fmt"
)

type Float64Col struct {
	rows []float64
}

func (c *Float64Col) Len() int {
	return len(c.rows)
}

func (c *Float64Col) Append(row interface{}) error {
	value, found := row.(float64)
	if !found {
		return errors.New(
			fmt.Sprintf("Wrong type (actual) %T != float64 (expected)", row),
		)
	}
	c.rows = append(c.rows, value)
	return nil
}
