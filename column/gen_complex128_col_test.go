// Test case for Complex128Col
// This file was generated. See generator directory.
package column

import "testing"

func TestComplex128ColAppend(t *testing.T) {
	c := new(Complex128Col)
	if err := c.Append(*new(complex128)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(uintptr)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestComplex128ColLen(t *testing.T) {
	c := new(Complex128Col)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(complex128))
	c.Append(*new(complex128))
	c.Append(*new(complex128))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}

func TestComplex128ColAt(t *testing.T) {
	c := new(Complex128Col)
	expected := *new(complex128)
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
