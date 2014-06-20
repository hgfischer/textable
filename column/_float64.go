package column

type Float64Col struct {
	rows []float64
}

func (c *Float64Col) Len() int {
	return len(c.rows)
}
