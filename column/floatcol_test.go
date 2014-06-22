package column

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var goodFloatConversions = map[interface{}]float64{
	float32(32.0):      float64(32.0),
	float64(6464.6464): float64(6464.6464),
	int(-123):          float64(-123),
	int8(-88):          float64(-88),
	int16(-1616):       float64(-1616),
	int32(-323232):     float64(-323232),
	int64(-64646464):   float64(-64646464),
	uint(123):          float64(123),
	uint8(88):          float64(88),
	uint16(1616):       float64(1616),
	uint32(323232):     float64(323232),
	uint64(64646464):   float64(64646464),
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

		// Reuse loop to test string conversion
		err = c.Append(fmt.Sprint(value))
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
