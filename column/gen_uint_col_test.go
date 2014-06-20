// Test case for UintCol
// This file was generated. See generator directory.
package column

import "testing"

func TestUintColAppend(t *testing.T) {
	c := new(UintCol)
	if err := c.Append(*new(uint)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(int64)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestUintColLen(t *testing.T) {
	c := new(UintCol)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(uint))
	c.Append(*new(uint))
	c.Append(*new(uint))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}
