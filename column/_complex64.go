package column

type Complex64Col struct {
	rows []complex64
}

func (c *Complex64Col) Len() int {
	return len(c.rows)
}
