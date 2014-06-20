package column

type RuneCol struct {
	rows []rune
}

func (c *RuneCol) Len() int {
	return len(c.rows)
}
