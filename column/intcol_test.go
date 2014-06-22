package column

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var goodIntConversions = map[interface{}]int64{
	float32(32.0):         32,
	float64(6464.6464):    6464,
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
	"31.1":                31,
	"-32.3":               -32,
	"31e1":                31,
	"-32E3":               -32,
	"31":                  31,
	"-32":                 -32,
	"0x10":                16,
	"0X20":                32,
	"071235":              29341,
	"-071235":             -29341,
	"085365":              85365,
	"-085365":             -85365,
	" 31.1    ":           31,
	" -32.3   ":           -32,
	" 31e1    ":           31,
	" -32E3   ":           -32,
	" 31      ":           31,
	" -32     ":           -32,
	" 0x10    ":           16,
	" 0X20    ":           32,
	" 071235  ":           29341,
	" -071235 ":           -29341,
	" 085365  ":           85365,
	" -085365 ":           -85365,
	" trash 31.1 trash  ": 31,
	" trash (-32.3)trash": -32,
	" trash 31e1trash   ": 31,
	" trash (-32E3)trash": -32,
	" trash 31trashsad  ": 31,
	" trash -32 cvbbnm d": -32,
	" trash 0x10s23ffass": 16,
	" trash 0X20vadasd w": 32,
	" trash 071235 fas s": 29341,
	" trash -071235 s dd": -29341,
	" trash 085365 s asd": 85365,
	" trash -085365 as d": -85365,
	"31.1, 151 1":         31,
	"-32.3) 6234":         -32,
	"31e1, 1  34":         31,
	"-32E3) 6234":         -32,
	"31 351.2 42":         31,
	"-32 354 433":         -32,
	"0x10 0x20 3":         16,
	"0X20 0123 4":         32,
	"071235 4222":         29341,
	"-071235 323":         -29341,
	"085365 0x20":         85365,
	"-085365 0x1":         -85365,
}

var badIntConversions = []interface{}{
	errors.New("some error"),
	"some random string",
}

func TestShouldAppendGoodConversionsToIntCol(t *testing.T) {
	var err error
	var at interface{}
	var ok bool

	c := new(IntCol)
	cnt := 0
	for value, expected := range goodIntConversions {
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

func TestShouldNotAppendBadConversionsToIntCol(t *testing.T) {
	var err error

	c := new(IntCol)
	for _, value := range badIntConversions {
		err = c.Append(value)
		assert.NotNil(t, err)       // Test successful conversion
		assert.Nil(t, c.Last())     // Test last value equality
		assert.Equal(t, 0, c.Len()) // Test Len()
	}
}

func TestShouldGetGoStringFromIntCol(t *testing.T) {
	c := new(IntCol)
	val := -622
	c.Append(val)
	c.Append(val)
	expected := fmt.Sprintf("[]int64{%#v, %#v}", val, val)
	value := c.String()
	if value != expected {
		t.Fatalf("Wrong value. (actual) %#v != %#v (expected)", value, expected)
	}
}
