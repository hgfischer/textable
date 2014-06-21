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

func (c *{{.CapsName}}) Len() uint {
	return uint(len(c.rows))
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

func (c *{{.CapsName}}) At(index uint) (value interface{}, exists bool) {
	if index > uint(len(c.rows)-1) {
		return nil, false
	}
	return c.rows[index], true
}

func (c *{{.CapsName}}) String() string {
	return fmt.Sprintf("%#v", c.rows)
}
