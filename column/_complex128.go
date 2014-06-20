package column

type Complex128Col struct {
	rows []complex128
}

func (c *Complex128Col) Len() int {
	return len(c.rows)
}
