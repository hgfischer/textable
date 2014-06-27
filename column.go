package textable

import "fmt"

type column struct {
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

func (c *column) Swap(i, j int) {
	c.values[i], c.values[j] = c.values[j], c.values[i]
	c.order[i], c.order[j] = c.order[j], c.order[i]
}

func (c *column) Less(i, j int) bool {
	// TODO properly detect type and make comparisons
	return c.values[i].(int) < c.values[j].(int)
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
