package column

type UintCol struct {
	rows []uint
}

func (c *UintCol) Len() int {
	return len(c.rows)
}
