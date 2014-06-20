// Test case for IntCol
// This file was generated. See generator directory.
package column

import "testing"

func TestIntColAppend(t *testing.T) {
	c := new(IntCol)
	if err := c.Append(*new(int)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(float64)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestIntColLen(t *testing.T) {
	c := new(IntCol)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(int))
	c.Append(*new(int))
	c.Append(*new(int))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}
