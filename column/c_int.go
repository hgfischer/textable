package column

type IntCol struct {
	rows []int
}

func (c *IntCol) Len() int {
	return len(c.rows)
}
