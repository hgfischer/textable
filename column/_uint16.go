package column

type Uint16Col struct {
	rows []uint16
}

func (c *Uint16Col) Len() int {
	return len(c.rows)
}
