package column

import "fmt"

type Column interface {
	Len() int
}

type StringerCol struct {
	rows []fmt.Stringer
}

func (c *StringerCol) Len() int {
	return len(c.rows)
}

type StringCol struct {
	rows []string
}

func (c *StringCol) Len() int {
	return len(c.rows)
}

type UInt8Col struct {
	rows []uint8
}

type UInt16Col struct {
	rows []uint16
}

type UInt32Col struct {
	rows []uint32
}

type UInt64Col struct {
	rows []uint64
}

type Int8Col struct {
	rows []uint8
}

type Int16Col struct {
	rows []uint16
}

type Int32Col struct {
	rows []uint32
}

type Int64Col struct {
	rows []uint64
}

type Float32Col struct {
	rows []float64
}

type Float64Col struct {
	rows []float64
}

type Complex64Col struct {
	rows []complex64
}

type Complex128Col struct {
	rows []complex128
}

type BoolCol struct {
	rows []bool
}

type ByteCol struct {
	rows []byte
}

type RuneCol struct {
	rows []rune
}

type UIntCol struct {
	rows []uint
}

type IntCol struct {
	rows []int
}
