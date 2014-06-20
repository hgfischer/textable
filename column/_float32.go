package column

type Float32Col struct {
	rows []float32
}

func (c *Float32Col) Len() int {
	return len(c.rows)
}
