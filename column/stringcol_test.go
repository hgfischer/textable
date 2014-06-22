package column

import (
	"fmt"
	"testing"
)

func TestStringColAppend(t *testing.T) {
	c := new(StringCol)
	c.Append(string("random string"))
	c.Append(float64(7.8))
	c.Append(int(-622))
	c.Append(uint(0xbeec38))

}

func TestStringColLen(t *testing.T) {
	c := new(StringCol)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(string("random string"))
	c.Append(string("random string"))
	c.Append(string("random string"))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}

func TestStringColAt(t *testing.T) {
	c := new(StringCol)

	expected := *new(string)

	c.Append(expected)
	value, exists := c.At(0)
	if !exists {
		t.Fatal("Value at index 0 should exist")
	}
	if value != expected {
		t.Fatalf("Wrong value. (actual) %#v != %#v (expected)", value, expected)
	}
	if value, exists = c.At(1); exists {
		t.Fatal("Value at index 1 should not exist")
	}
}

func TestStringColString(t *testing.T) {
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
