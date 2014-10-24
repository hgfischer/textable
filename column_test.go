package textable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColumnValueAppendAndLen(t *testing.T) {
	values := []interface{}{1, 2.1, -3, -4.5, "string", "~", nil}
	c := new(column)
	assert.Equal(t, 0, c.Len())
	for k, value := range values {
		c.Append(value)
		assert.Equal(t, k+1, c.Len())
	}
}

func TestColumnValueAndOrderSwap(t *testing.T) {
	c := new(column)
	first, last := "a", "b"
	c.Append(first)
	c.Append(last)
	assert.Equal(t, first, c.values[0])
	assert.Equal(t, last, c.values[1])
	assert.Equal(t, 1, c.order[0])
	assert.Equal(t, 2, c.order[1])
	c.Swap(0, 1)
	assert.Equal(t, last, c.values[0])
	assert.Equal(t, first, c.values[1])
	assert.Equal(t, 2, c.order[0])
	assert.Equal(t, 1, c.order[1])
}

func TestColumnValueFormat(t *testing.T) {
	c := new(column)
	c.Append(1.2)
	c.Append("1.3")
	c.Append(nil)
	c.Append("string")
	c.SetFormat("%s")
	assert.Equal(t, len("string"), c.Width())
}
