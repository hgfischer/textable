package column

import (
	"fmt"
	"math"

	"github.com/hgfischer/textable/strng"
)

type UIntCol struct {
	rows []uint64
}

func (c *UIntCol) Len() uint {
	return uint(len(c.rows))
}

func (c *UIntCol) Append(row interface{}) error {
	switch value := row.(type) {
	case float32:
		c.rows = append(c.rows, uint64(math.Abs(float64(value))))
	case float64:
		c.rows = append(c.rows, uint64(math.Abs(float64(value))))
	case int:
		c.rows = append(c.rows, uint64(math.Abs(float64(value))))
	case int8:
		c.rows = append(c.rows, uint64(math.Abs(float64(value))))
	case int16:
		c.rows = append(c.rows, uint64(math.Abs(float64(value))))
	case int32:
		c.rows = append(c.rows, uint64(math.Abs(float64(value))))
	case int64:
		c.rows = append(c.rows, uint64(math.Abs(float64(value))))
	case uint:
		c.rows = append(c.rows, uint64(value))
	case uint8:
		c.rows = append(c.rows, uint64(value))
	case uint16:
		c.rows = append(c.rows, uint64(value))
	case uint32:
		c.rows = append(c.rows, uint64(value))
	case uint64:
		c.rows = append(c.rows, uint64(value))
	default:
		num, err := strng.ExtractUint(fmt.Sprint(value))
		if err != nil {
			return err
		}
		c.rows = append(c.rows, num)
	}
	return nil
}

func (c *UIntCol) At(index uint) (value interface{}, exists bool) {
	if index < c.Len() {
		return c.rows[index], true
	}
	return nil, false
}

func (c *UIntCol) String() string {
	return fmt.Sprintf("%#v", c.rows)
}

func (c *UIntCol) Last() (value interface{}) {
	if pos := c.Len(); pos != 0 {
		value, _ = c.At(pos - 1)
	}
	return
}
