package column

import (
	"fmt"
	"github.com/hgfischer/textable/strng"
)

type FloatCol struct {
	rows []float64
}

func (c *FloatCol) Len() uint {
	return uint(len(c.rows))
}

func (c *FloatCol) Append(row interface{}) error {
	switch value := row.(type) {
	case float32:
		c.rows = append(c.rows, float64(value))
	case float64:
		c.rows = append(c.rows, float64(value))
	case int:
		c.rows = append(c.rows, float64(value))
	case int8:
		c.rows = append(c.rows, float64(value))
	case int16:
		c.rows = append(c.rows, float64(value))
	case int32:
		c.rows = append(c.rows, float64(value))
	case int64:
		c.rows = append(c.rows, float64(value))
	case uint:
		c.rows = append(c.rows, float64(value))
	case uint8:
		c.rows = append(c.rows, float64(value))
	case uint16:
		c.rows = append(c.rows, float64(value))
	case uint32:
		c.rows = append(c.rows, float64(value))
	case uint64:
		c.rows = append(c.rows, float64(value))
	default:
		num, err := strng.ExtractFloat(fmt.Sprint(value))
		if err != nil {
			return err
		}
		c.rows = append(c.rows, num)
	}
	return nil
}

func (c *FloatCol) At(index uint) (value interface{}, exists bool) {
	if index < c.Len() {
		return c.rows[index], true
	}
	return nil, false
}

func (c *FloatCol) String() string {
	return fmt.Sprintf("%#v", c.rows)
}

func (c *FloatCol) Last() (value interface{}) {
	if pos := c.Len(); pos != 0 {
		value, _ = c.At(pos - 1)
	}
	return
}
