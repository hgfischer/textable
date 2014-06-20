package column

type Int64Col struct {
	rows []int64
}

func (c *Int64Col) Len() int {
	return len(c.rows)
}
