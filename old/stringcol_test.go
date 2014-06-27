package column

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var goodStringConversions = map[interface{}]string{
	string(" trailing spaces "): "trailing spaces",
	string("random string"):     "random string",
	float64(7.8):                "7.8",
	int(-622):                   "-622",
	uint(0xbeec38):              "12512312",
	nil:                         "",
}

func TestShouldAppendGoodConversionsToStringCol(t *testing.T) {
	var err error
	var at interface{}
	var ok bool

	c := new(StringCol)
	cnt := 0
	for value, expected := range goodStringConversions {
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

func TestShouldGetGoStringFromStringCol(t *testing.T) {
	c := new(StringCol)
	val := string("random string")
	c.Append(val)
	c.Append(val)
	expected := fmt.Sprintf("[]string{%#v, %#v}", val, val)
	value := c.String()
	if value != expected {
		t.Fatalf("Wrong value. (actual) %#v != %#v (expected)", value, expected)
	}
}

func TestShouldReturnNilFalseIfIndexDoesNotExistInStringCol(t *testing.T) {
	c := new(StringCol)
	v, ok := c.At(10)
	assert.Nil(t, v)
	assert.False(t, ok)
}
