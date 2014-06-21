// Test case for Float64Col
// This file was generated. See generator directory.
package column

import "testing"

func TestFloat64ColAppend(t *testing.T) {
	c := new(Float64Col)
	if err := c.Append(*new(float64)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(float32)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestFloat64ColLen(t *testing.T) {
	c := new(Float64Col)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(float64))
	c.Append(*new(float64))
	c.Append(*new(float64))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}

func TestFloat64ColAt(t *testing.T) {
	c := new(Float64Col)
	expected := *new(float64)
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
