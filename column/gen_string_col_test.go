// Test case for StringCol
// This file was generated. See generator directory.
package column

import "testing"

func TestStringColAppend(t *testing.T) {
	c := new(StringCol)
	if err := c.Append(*new(string)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(uint64)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestStringColLen(t *testing.T) {
	c := new(StringCol)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(string))
	c.Append(*new(string))
	c.Append(*new(string))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}

func TestStringColAt(t *testing.T) {
	c := new(StringCol)
	expected := *new(string)
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
