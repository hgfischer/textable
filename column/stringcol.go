package column

import "fmt"

type StringCol struct {
	rows []string
}

func (c *StringCol) Len() uint {
	return uint(len(c.rows))
}

func (c *StringCol) Append(row interface{}) error {
	c.rows = append(c.rows, fmt.Sprint(row))
	return nil
}

func (c *StringCol) At(index uint) (value interface{}, exists bool) {
	if index > uint(len(c.rows)-1) {
		return nil, false
	}
	return c.rows[index], true
}

func (c *StringCol) String() string {
	return fmt.Sprintf("%#v", c.rows)
}

func (c *StringCol) Last() (value interface{}) {
	if pos := c.Len(); pos != 0 {
		value, _ = c.At(pos)
	}
	return
}
