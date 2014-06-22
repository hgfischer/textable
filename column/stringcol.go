package column

import (
	"fmt"
	"strings"
)

type StringCol struct {
	rows []string
}

func (c *StringCol) Len() uint {
	return uint(len(c.rows))
}

func (c *StringCol) Append(row interface{}) error {
	str := ""
	if row != nil {
		str = strings.TrimSpace(fmt.Sprint(row))
	}
	c.rows = append(c.rows, str)
	return nil
}

func (c *StringCol) At(index uint) (value interface{}, exists bool) {
	if index < c.Len() {
		return c.rows[index], true
	}
	return nil, false
}

func (c *StringCol) String() string {
	return fmt.Sprintf("%#v", c.rows)
}

func (c *StringCol) Last() (value interface{}) {
	if pos := c.Len(); pos != 0 {
		value, _ = c.At(pos - 1)
	}
	return
}
