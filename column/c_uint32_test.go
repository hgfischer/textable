// Test case for Uint32Col
// This file was generated. See generator directory.
package column

import "testing"

func TestUint32ColAppend(t *testing.T) {
	c := new(Uint32Col)
	if err := c.Append(*new(uint32)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(uint16)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestUint32ColLen(t *testing.T) {
	c := new(Uint32Col)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(uint32))
	c.Append(*new(uint32))
	c.Append(*new(uint32))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}
