package textable

import (
	"fmt"
	"reflect"
)

type column struct {
	kind          interface{}
	align         Align
	order         []int
	values        []interface{}
	valuesChanged bool
	fmtdValues    []string
	formatChanged bool
	format        string
	width         uint
}

func (c *column) Append(value interface{}) {
	c.values = append(c.values, value)
	c.order = append(c.order, len(c.values))
	c.valuesChanged = true
}

func (c *column) Len() int {
	return len(c.values)
}

func (c *column) Swap(a, b int) {
	c.values[a], c.values[b] = c.values[b], c.values[a]
	c.order[a], c.order[b] = c.order[b], c.order[a]
}

func (c *column) Less(a, b int) bool {
	valA := reflect.ValueOf(c.values[a])

	valB := reflect.ValueOf(c.values[a])

	return c.values[a].(int) < c.values[b].(int)
}

func (c *column) SetFormat(format string) error {
	c.format = format
	c.formatChanged = true
	// TODO check if format is compatible with type
	return nil
}

func (c *column) renderFormattedRows() {
	if len(c.fmtdValues) != len(c.values) {
		c.fmtdValues = make([]string, len(c.values))
	}
	for k, r := range c.values {
		c.fmtdValues[k] = fmt.Sprintf(c.format, r)
	}
	fmt.Println(c.fmtdValues)
}

func (c *column) Width() uint {
	if c.valuesChanged || c.formatChanged {
		c.renderFormattedRows()
		c.valuesChanged = false
		c.formatChanged = false
	}
	var max int
	for _, r := range c.fmtdValues {
		curr := len(r)
		if curr > max {
			max = curr
		}
	}
	return uint(max)
}
