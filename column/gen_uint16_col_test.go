// Test case for Uint16Col
// This file was generated. See generator directory.
package column

import (
	"fmt"
	"testing"
)

func TestUint16ColAppend(t *testing.T) {
	c := new(Uint16Col)
	if err := c.Append(*new(uint16)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(byte)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestUint16ColLen(t *testing.T) {
	c := new(Uint16Col)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(uint16))
	c.Append(*new(uint16))
	c.Append(*new(uint16))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}

func TestUint16ColAt(t *testing.T) {
	c := new(Uint16Col)
	expected := *new(uint16)
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

func TestUint16ColString(t *testing.T) {
	c := new(Uint16Col)
	val := *new(uint16)
	c.Append(val)
	c.Append(val)
	expected := fmt.Sprintf("[]uint16{%#v, %#v}", val, val)
	value := c.String()
	if value != expected {
		t.Fatalf("Wrong value. (actual) %#v != %#v (expected)", value, expected)
	}
}
