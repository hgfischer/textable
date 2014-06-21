// Test case for BoolCol
// This file was generated. See generator directory.
package column

import "testing"

func TestBoolColAppend(t *testing.T) {
	c := new(BoolCol)
	if err := c.Append(*new(bool)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(string)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestBoolColLen(t *testing.T) {
	c := new(BoolCol)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(bool))
	c.Append(*new(bool))
	c.Append(*new(bool))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}

func TestBoolColAt(t *testing.T) {
	c := new(BoolCol)
	expected := *new(bool)
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
