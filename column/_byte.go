package column

type ByteCol struct {
	rows []byte
}

func (c *ByteCol) Len() int {
	return len(c.rows)
}
