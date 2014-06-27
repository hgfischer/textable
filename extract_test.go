package textable

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var goodFloatConversions = map[string]float64{
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

var goodIntConversions = map[string]int64{
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

var goodUintConversions = map[string]uint64{
	"31.1":                31,
	"-32.3":               32,
	"31e1":                31,
	"-32E3":               32,
	"31":                  31,
	"-32":                 32,
	"0x10":                16,
	"0X20":                32,
	"071235":              29341,
	"-071235":             29341,
	"085365":              85365,
	"-085365":             85365,
	" 31.1    ":           31,
	" -32.3   ":           32,
	" 31e1    ":           31,
	" -32E3   ":           32,
	" 31      ":           31,
	" -32     ":           32,
	" 0x10    ":           16,
	" 0X20    ":           32,
	" 071235  ":           29341,
	" -071235 ":           29341,
	" 085365  ":           85365,
	" -085365 ":           85365,
	" trash 31.1 trash  ": 31,
	" trash (-32.3)trash": 32,
	" trash 31e1trash   ": 31,
	" trash (-32E3)trash": 32,
	" trash 31trashsad  ": 31,
	" trash -32 cvbbnm d": 32,
	" trash 0x10s23ffass": 16,
	" trash 0X20vadasd w": 32,
	" trash 071235 fas s": 29341,
	" trash -071235 s dd": 29341,
	" trash 085365 s asd": 85365,
	" trash -085365 as d": 85365,
	"31.1, 151 1":         31,
	"-32.3) 6234":         32,
	"31e1, 1  34":         31,
	"-32E3) 6234":         32,
	"31 351.2 42":         31,
	"-32 354 433":         32,
	"0x10 0x20 3":         16,
	"0X20 0123 4":         32,
	"071235 4222":         29341,
	"-071235 323":         29341,
	"085365 0x20":         85365,
	"-085365 0x1":         85365,
}

var badConversions = []string{
	"some random string",
}

func TestFloatConversions(t *testing.T) {
	for input, expected := range goodFloatConversions {
		converted, err := ExtractFloat(input)
		assert.Nil(t, err)
		assert.Equal(t, expected, converted,
			fmt.Sprintf("(input) %#v => %#v[%[2]f] (expected)", input, expected))
	}
	for _, input := range badConversions {
		converted, err := ExtractFloat(input)
		assert.NotNil(t, err, input)
		assert.Empty(t, converted, input)
	}
}

func TestIntConversions(t *testing.T) {
	for input, expected := range goodIntConversions {
		converted, err := ExtractInt(input)
		assert.Nil(t, err)
		assert.Equal(t, expected, converted,
			fmt.Sprintf("(input) %#v => %#v[%[2]d] (expected)", input, expected))
	}
	for _, input := range badConversions {
		converted, err := ExtractInt(input)
		assert.NotNil(t, err, input)
		assert.Empty(t, converted, input)
	}
}

func TestUintConversions(t *testing.T) {
	for input, expected := range goodUintConversions {
		converted, err := ExtractUint(input)
		assert.Nil(t, err)
		assert.Equal(t, expected, converted,
			fmt.Sprintf("(input) %#v => %#v[%[2]d] (expected)", input, expected))
	}
	for _, input := range badConversions {
		converted, err := ExtractUint(input)
		assert.NotNil(t, err, input)
		assert.Empty(t, converted, input)
	}
}
