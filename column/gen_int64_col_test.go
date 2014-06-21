// Test case for Int64Col
// This file was generated. See generator directory.
package column

import (
	"fmt"
	"testing"
)

func TestInt64ColAppend(t *testing.T) {
	c := new(Int64Col)
	if err := c.Append(*new(int64)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(int32)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestInt64ColLen(t *testing.T) {
	c := new(Int64Col)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(int64))
	c.Append(*new(int64))
	c.Append(*new(int64))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}

func TestInt64ColAt(t *testing.T) {
	c := new(Int64Col)
	expected := *new(int64)
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

func TestInt64ColString(t *testing.T) {
	c := new(Int64Col)
	val := *new(int64)
	c.Append(val)
	c.Append(val)
	expected := fmt.Sprintf("[]int64{%#v, %#v}", val, val)
	value := c.String()
	if value != expected {
		t.Fatalf("Wrong value. (actual) %#v != %#v (expected)", value, expected)
	}
}
