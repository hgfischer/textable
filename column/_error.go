package column

type ErrorCol struct {
	rows []error
}

func (c *ErrorCol) Len() int {
	return len(c.rows)
}
