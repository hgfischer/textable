// Test case for Int8Col
// This file was generated. See generator directory.
package column

import (
	"fmt"
	"testing"
)

func TestInt8ColAppend(t *testing.T) {
	c := new(Int8Col)
	if err := c.Append(*new(int8)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(int)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestInt8ColLen(t *testing.T) {
	c := new(Int8Col)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(int8))
	c.Append(*new(int8))
	c.Append(*new(int8))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}

func TestInt8ColAt(t *testing.T) {
	c := new(Int8Col)
	expected := *new(int8)
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

func TestInt8ColString(t *testing.T) {
	c := new(Int8Col)
	val := *new(int8)
	c.Append(val)
	c.Append(val)
	expected := fmt.Sprintf("[]int8{%#v, %#v}", val, val)
	value := c.String()
	if value != expected {
		t.Fatalf("Wrong value. (actual) %#v != %#v (expected)", value, expected)
	}
}
