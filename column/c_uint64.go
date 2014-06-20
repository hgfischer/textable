package column

type Uint64Col struct {
	rows []uint64
}

func (c *Uint64Col) Len() int {
	return len(c.rows)
}
