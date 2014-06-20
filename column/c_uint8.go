package column

type Uint8Col struct {
	rows []uint8
}

func (c *Uint8Col) Len() int {
	return len(c.rows)
}
