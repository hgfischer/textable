package column

type Int8Col struct {
	rows []int8
}

func (c *Int8Col) Len() int {
	return len(c.rows)
}
