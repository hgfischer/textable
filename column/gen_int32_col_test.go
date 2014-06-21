// Test case for Int32Col
// This file was generated. See generator directory.
package column

import (
	"fmt"
	"testing"
)

func TestInt32ColAppend(t *testing.T) {
	c := new(Int32Col)
	if err := c.Append(*new(int32)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(int16)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestInt32ColLen(t *testing.T) {
	c := new(Int32Col)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(int32))
	c.Append(*new(int32))
	c.Append(*new(int32))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}

func TestInt32ColAt(t *testing.T) {
	c := new(Int32Col)
	expected := *new(int32)
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

func TestInt32ColString(t *testing.T) {
	c := new(Int32Col)
	val := *new(int32)
	c.Append(val)
	c.Append(val)
	expected := fmt.Sprintf("[]int32{%#v, %#v}", val, val)
	value := c.String()
	if value != expected {
		t.Fatalf("Wrong value. (actual) %#v != %#v (expected)", value, expected)
	}
}
