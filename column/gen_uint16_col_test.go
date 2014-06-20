// Test case for Uint16Col
// This file was generated. See generator directory.
package column

import "testing"

func TestUint16ColAppend(t *testing.T) {
	c := new(Uint16Col)
	if err := c.Append(*new(uint16)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(uint8)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestUint16ColLen(t *testing.T) {
	c := new(Uint16Col)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(uint16))
	c.Append(*new(uint16))
	c.Append(*new(uint16))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}
