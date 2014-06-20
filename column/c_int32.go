package column

type Int32Col struct {
	rows []int32
}

func (c *Int32Col) Len() int {
	return len(c.rows)
}
