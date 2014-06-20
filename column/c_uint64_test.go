// Test case for Uint64Col
// This file was generated. See generator directory.
package column

import "testing"

func TestUint64ColAppend(t *testing.T) {
	c := new(Uint64Col)
	if err := c.Append(*new(uint64)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(uint32)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestUint64ColLen(t *testing.T) {
	c := new(Uint64Col)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(uint64))
	c.Append(*new(uint64))
	c.Append(*new(uint64))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}
