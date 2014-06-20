// Float32Col
// This file was generated. See generator directory.
package column

import (
	"errors"
	"fmt"
)

type Float32Col struct {
	rows []float32
}

func (c *Float32Col) Len() int {
	return len(c.rows)
}

func (c *Float32Col) Append(row interface{}) error {
	value, found := row.(float32)
	if !found {
		return errors.New(
			fmt.Sprintf("Wrong type (actual) %T != float32 (expected)", row),
		)
	}
	c.rows = append(c.rows, value)
	return nil
}
