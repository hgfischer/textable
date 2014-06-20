// {{.CapsName}}
// This file was generated. See generator directory.
package column

import (
	"errors"
	"fmt"
)

type {{.CapsName}} struct {
	rows []{{.Type}}
}

func (c *{{.CapsName}}) Len() int {
	return len(c.rows)
}

func (c *{{.CapsName}}) Append(row interface{}) error {
	value, found := row.({{.Type}})
	if !found {
		return errors.New(
			fmt.Sprintf("Wrong type (actual) %T != {{.Type}} (expected)", row),
		)
	}
	c.rows = append(c.rows, value)
	return nil
}
