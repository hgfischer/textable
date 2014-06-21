// Test case for Uint64Col
// This file was generated. See generator directory.
package column

import (
	"fmt"
	"testing"
)

func TestUint64ColAppend(t *testing.T) {
	c := new(Uint64Col)
	if err := c.Append(*new(uint64)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(uint32)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestUint64ColLen(t *testing.T) {
	c := new(Uint64Col)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(uint64))
	c.Append(*new(uint64))
	c.Append(*new(uint64))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}

func TestUint64ColAt(t *testing.T) {
	c := new(Uint64Col)
	expected := *new(uint64)
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

func TestUint64ColString(t *testing.T) {
	c := new(Uint64Col)
	val := *new(uint64)
	c.Append(val)
	c.Append(val)
	expected := fmt.Sprintf("[]uint64{%#v, %#v}", val, val)
	value := c.String()
	if value != expected {
		t.Fatalf("Wrong value. (actual) %#v != %#v (expected)", value, expected)
	}
}
