package column

type {{.CapsName}}Col struct {
	rows []{{.Type}}
}

func (c *{{.CapsName}}Col) Len() int {
	return len(c.rows)
}
