package column

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SomeStringer struct {
	Value string
}

func (ss *SomeStringer) String() string {
	return ss.Value
}

func TestNewForTypeOfShouldGetAStringCol(t *testing.T) {
	var col Column
	var ok bool
	inputs := []interface{}{"string", SomeStringer{"string"}}
	for _, input := range inputs {
		col = NewForTypeOf(input)
		_, ok = col.(*StringCol)
		assert.True(t, ok, fmt.Sprint(input))
		assert.Nil(t, col.Last()) // Should not append the value
	}
}

func TestNewForTypeOfShouldGetAFloatCol(t *testing.T) {
	var col Column
	var ok bool
	inputs := []interface{}{float64(1.2), float32(3.4)}
	for _, input := range inputs {
		col = NewForTypeOf(input)
		_, ok = col.(*FloatCol)
		assert.True(t, ok, fmt.Sprint(input))
		assert.Nil(t, col.Last()) // Should not append the value
	}
}

func TestNewForTypeOfShouldGetAIntCol(t *testing.T) {
	var col Column
	var ok bool
	inputs := []interface{}{int(-1), int8(-8), int16(-16), int32(-32), int64(-64)}
	for _, input := range inputs {
		col = NewForTypeOf(input)
		_, ok = col.(*IntCol)
		assert.True(t, ok, fmt.Sprint(input))
		assert.Nil(t, col.Last()) // Should not append the value
	}
}

func TestNewForTypeOfShouldGetAUIntCol(t *testing.T) {
	var col Column
	var ok bool
	inputs := []interface{}{uint(1), uint8(8), uint16(16), uint32(32), uint64(64)}
	for _, input := range inputs {
		col = NewForTypeOf(input)
		_, ok = col.(*UIntCol)
		assert.True(t, ok, fmt.Sprint(input))
		assert.Nil(t, col.Last()) // Should not append the value
	}
}
