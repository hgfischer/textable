// Test case for UintCol
// This file was generated. See generator directory.
package column

import (
	"fmt"
	"testing"
)

func TestUintColAppend(t *testing.T) {
	c := new(UintCol)
	if err := c.Append(*new(uint)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(int64)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestUintColLen(t *testing.T) {
	c := new(UintCol)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(uint))
	c.Append(*new(uint))
	c.Append(*new(uint))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}

func TestUintColAt(t *testing.T) {
	c := new(UintCol)
	expected := *new(uint)
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

func TestUintColString(t *testing.T) {
	c := new(UintCol)
	val := *new(uint)
	c.Append(val)
	c.Append(val)
	expected := fmt.Sprintf("[]uint{%#v, %#v}", val, val)
	value := c.String()
	if value != expected {
		t.Fatalf("Wrong value. (actual) %#v != %#v (expected)", value, expected)
	}
}
