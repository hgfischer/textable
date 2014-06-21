// Test case for IntCol
// This file was generated. See generator directory.
package column

import (
	"fmt"
	"testing"
)

func TestIntColAppend(t *testing.T) {
	c := new(IntCol)
	if err := c.Append(*new(int)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(float64)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestIntColLen(t *testing.T) {
	c := new(IntCol)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(int))
	c.Append(*new(int))
	c.Append(*new(int))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}

func TestIntColAt(t *testing.T) {
	c := new(IntCol)
	expected := *new(int)
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

func TestIntColString(t *testing.T) {
	c := new(IntCol)
	val := *new(int)
	c.Append(val)
	c.Append(val)
	expected := fmt.Sprintf("[]int{%#v, %#v}", val, val)
	value := c.String()
	if value != expected {
		t.Fatalf("Wrong value. (actual) %#v != %#v (expected)", value, expected)
	}
}
