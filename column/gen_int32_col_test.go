// Test case for Int32Col
// This file was generated. See generator directory.
package column

import "testing"

func TestInt32ColAppend(t *testing.T) {
	c := new(Int32Col)
	if err := c.Append(*new(int32)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(int16)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestInt32ColLen(t *testing.T) {
	c := new(Int32Col)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(int32))
	c.Append(*new(int32))
	c.Append(*new(int32))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}
