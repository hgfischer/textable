// StringCol
// This file was generated. See generator directory.
package column

import (
	"errors"
	"fmt"
)

type StringCol struct {
	rows []string
}

func (c *StringCol) Len() int {
	return len(c.rows)
}

func (c *StringCol) Append(row interface{}) error {
	value, found := row.(string)
	if !found {
		return errors.New(
			fmt.Sprintf("Wrong type (actual) %T != string (expected)", row),
		)
	}
	c.rows = append(c.rows, value)
	return nil
}
