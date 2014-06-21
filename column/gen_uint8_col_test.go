// Test case for Uint8Col
// This file was generated. See generator directory.
package column

import "testing"

func TestUint8ColAppend(t *testing.T) {
	c := new(Uint8Col)
	if err := c.Append(*new(uint8)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(uint)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestUint8ColLen(t *testing.T) {
	c := new(Uint8Col)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(uint8))
	c.Append(*new(uint8))
	c.Append(*new(uint8))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}

func TestUint8ColAt(t *testing.T) {
	c := new(Uint8Col)
	expected := *new(uint8)
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
