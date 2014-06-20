// Test case for RuneCol
// This file was generated. See generator directory.
package column

import "testing"

func TestRuneColAppend(t *testing.T) {
	c := new(RuneCol)
	if err := c.Append(*new(rune)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(byte)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestRuneColLen(t *testing.T) {
	c := new(RuneCol)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(rune))
	c.Append(*new(rune))
	c.Append(*new(rune))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}
