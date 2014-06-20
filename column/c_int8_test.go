// Test case for Int8Col
// This file was generated. See generator directory.
package column

import "testing"

func TestInt8ColAppend(t *testing.T) {
	c := new(Int8Col)
	if err := c.Append(*new(int8)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(int)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestInt8ColLen(t *testing.T) {
	c := new(Int8Col)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(int8))
	c.Append(*new(int8))
	c.Append(*new(int8))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}
