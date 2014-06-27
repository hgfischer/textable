package column

import (
	"fmt"

	"github.com/hgfischer/textable/strng"
)

type IntCol struct {
	rows []int64
}

func (c *IntCol) Len() uint {
	return uint(len(c.rows))
}

func (c *IntCol) Append(row interface{}) error {
	switch value := row.(type) {
	case float32:
		c.rows = append(c.rows, int64(value))
	case float64:
		c.rows = append(c.rows, int64(value))
	case int:
		c.rows = append(c.rows, int64(value))
	case int8:
		c.rows = append(c.rows, int64(value))
	case int16:
		c.rows = append(c.rows, int64(value))
	case int32:
		c.rows = append(c.rows, int64(value))
	case int64:
		c.rows = append(c.rows, int64(value))
	case uint:
		c.rows = append(c.rows, int64(value))
	case uint8:
		c.rows = append(c.rows, int64(value))
	case uint16:
		c.rows = append(c.rows, int64(value))
	case uint32:
		c.rows = append(c.rows, int64(value))
	case uint64:
		c.rows = append(c.rows, int64(value))
	default:
		num, err := strng.ExtractInt(fmt.Sprint(value))
		if err != nil {
			return err
		}
		c.rows = append(c.rows, num)
	}
	return nil
}

func (c *IntCol) At(index uint) (value interface{}, exists bool) {
	if index < c.Len() {
		return c.rows[index], true
	}
	return nil, false
}

func (c *IntCol) String() string {
	return fmt.Sprintf("%#v", c.rows)
}

func (c *IntCol) Last() (value interface{}) {
	if pos := c.Len(); pos != 0 {
		value, _ = c.At(pos - 1)
	}
	return
}
