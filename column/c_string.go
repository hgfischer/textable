package column

type StringCol struct {
	rows []string
}

func (c *StringCol) Len() int {
	return len(c.rows)
}
