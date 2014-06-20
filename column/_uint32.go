package column

type Uint32Col struct {
	rows []uint32
}

func (c *Uint32Col) Len() int {
	return len(c.rows)
}
