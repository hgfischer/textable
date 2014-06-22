package column

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var goodFloatConversions = map[interface{}]float64{
	float32(32.0):         32.0,
	float64(6464.6464):    6464.6464,
	int(-123):             -123,
	int8(-88):             -88,
	int16(-1616):          -1616,
	int32(-323232):        -323232,
	int64(-64646464):      -64646464,
	uint(123):             123,
	uint8(88):             88,
	uint16(1616):          1616,
	uint32(323232):        323232,
	uint64(64646464):      64646464,
	"31.1":                31.1,
	"-32.3":               -32.3,
	"31e1":                31e1,
	"-32E3":               -32e3,
	" 31.1    ":           31.1,
	" -32.3   ":           -32.3,
	" 31e1    ":           31e1,
	" -32E3   ":           -32e3,
	" trash 31.1 trash  ": 31.1,
	" trash (-32.3)trash": -32.3,
	" trash 31e1trash   ": 31e1,
	" trash (-32E3)trash": -32e3,
	"31.1, 151 1":         31.1,
	"-32.3) 6234":         -32.3,
	"31e1, 1  34":         31e1,
	"-32E3) 6234":         -32e3,
	"31 351.2 42":         31,
	"-32 354 433":         -32,
	"0x10 0x20 3":         0,
	"071235 4222":         71235,
	"-071235 323":         -71235,
}

var badFloatConversions = []interface{}{
	errors.New("some error"),
	"some random string",
}

func TestShouldAppendGoodConversionsToFloatCol(t *testing.T) {
	var err error
	var at interface{}
	var ok bool

	c := new(FloatCol)
	cnt := 0
	for value, expected := range goodFloatConversions {
		err = c.Append(value)
		cnt = cnt + 1
		assert.Nil(t, err)                  // Test successful conversion
		assert.Equal(t, expected, c.Last()) // Test last value equality
		assert.Equal(t, cnt, c.Len())       // Test Len()
		at, ok = c.At(uint(cnt - 1))        // Get last value with At() method and check again
		assert.True(t, ok)
		assert.Equal(t, expected, at)
	}
}

func TestShouldNotAppendBadConversionsToFloatCol(t *testing.T) {
	var err error

	c := new(FloatCol)
	for _, value := range badFloatConversions {
		err = c.Append(value)
		assert.NotNil(t, err)       // Test successful conversion
		assert.Nil(t, c.Last())     // Test last value equality
		assert.Equal(t, 0, c.Len()) // Test Len()
	}
}

func TestShouldGetGoStringFromFloatCol(t *testing.T) {
	c := new(FloatCol)
	val := float64(7.8)
	c.Append(val)
	c.Append(val)
	expected := fmt.Sprintf("[]float64{%#v, %#v}", val, val)
	value := c.String()
	if value != expected {
		t.Fatalf("Wrong value. (actual) %#v != %#v (expected)", value, expected)
	}
}

func TestShouldReturnNilFalseIfIndexDoesNotExistInFloatCol(t *testing.T) {
	c := new(FloatCol)
	v, ok := c.At(10)
	assert.Nil(t, v)
	assert.False(t, ok)
}
