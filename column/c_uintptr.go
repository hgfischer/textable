package column

type UintptrCol struct {
	rows []uintptr
}

func (c *UintptrCol) Len() int {
	return len(c.rows)
}
