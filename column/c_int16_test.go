// Test case for Int16Col
// This file was generated. See generator directory.
package column

import "testing"

func TestInt16ColAppend(t *testing.T) {
	c := new(Int16Col)
	if err := c.Append(*new(int16)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(int8)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestInt16ColLen(t *testing.T) {
	c := new(Int16Col)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(int16))
	c.Append(*new(int16))
	c.Append(*new(int16))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}
