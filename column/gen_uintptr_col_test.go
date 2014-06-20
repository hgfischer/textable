// Test case for UintptrCol
// This file was generated. See generator directory.
package column

import "testing"

func TestUintptrColAppend(t *testing.T) {
	c := new(UintptrCol)
	if err := c.Append(*new(uintptr)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(bool)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestUintptrColLen(t *testing.T) {
	c := new(UintptrCol)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(uintptr))
	c.Append(*new(uintptr))
	c.Append(*new(uintptr))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}
