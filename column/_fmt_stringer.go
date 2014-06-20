package column

type StringerCol struct {
	rows []fmt.Stringer
}

func (c *StringerCol) Len() int {
	return len(c.rows)
}
