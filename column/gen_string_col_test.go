// Test case for StringCol
// This file was generated. See generator directory.
package column

import "testing"

func TestStringColAppend(t *testing.T) {
	c := new(StringCol)
	if err := c.Append(*new(string)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(uint64)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestStringColLen(t *testing.T) {
	c := new(StringCol)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(string))
	c.Append(*new(string))
	c.Append(*new(string))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}
