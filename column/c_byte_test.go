// Test case for ByteCol
// This file was generated. See generator directory.
package column

import "testing"

func TestByteColAppend(t *testing.T) {
	c := new(ByteCol)
	if err := c.Append(*new(byte)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(bool)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestByteColLen(t *testing.T) {
	c := new(ByteCol)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(byte))
	c.Append(*new(byte))
	c.Append(*new(byte))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}
