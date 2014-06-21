// Test case for Float32Col
// This file was generated. See generator directory.
package column

import (
	"fmt"
	"testing"
)

func TestFloat32ColAppend(t *testing.T) {
	c := new(Float32Col)
	if err := c.Append(*new(float32)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(complex64)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestFloat32ColLen(t *testing.T) {
	c := new(Float32Col)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(float32))
	c.Append(*new(float32))
	c.Append(*new(float32))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}

func TestFloat32ColAt(t *testing.T) {
	c := new(Float32Col)
	expected := *new(float32)
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

func TestFloat32ColString(t *testing.T) {
	c := new(Float32Col)
	val := *new(float32)
	c.Append(val)
	c.Append(val)
	expected := fmt.Sprintf("[]float32{%#v, %#v}", val, val)
	value := c.String()
	if value != expected {
		t.Fatalf("Wrong value. (actual) %#v != %#v (expected)", value, expected)
	}
}
