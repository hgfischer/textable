package column

type BoolCol struct {
	rows []bool
}

func (c *BoolCol) Len() int {
	return len(c.rows)
}
