// Test case for BoolCol
// This file was generated. See generator directory.
package column

import "testing"

func TestBoolColAppend(t *testing.T) {
	c := new(BoolCol)
	if err := c.Append(*new(bool)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(uint64)); err == nil {
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
