// Test case for Int64Col
// This file was generated. See generator directory.
package column

import "testing"

func TestInt64ColAppend(t *testing.T) {
	c := new(Int64Col)
	if err := c.Append(*new(int64)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(int32)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestInt64ColLen(t *testing.T) {
	c := new(Int64Col)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(int64))
	c.Append(*new(int64))
	c.Append(*new(int64))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}
