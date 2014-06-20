package column

type Int16Col struct {
	rows []int16
}

func (c *Int16Col) Len() int {
	return len(c.rows)
}
