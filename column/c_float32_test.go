// Test case for Float32Col
// This file was generated. See generator directory.
package column

import "testing"

func TestFloat32ColAppend(t *testing.T) {
	c := new(Float32Col)
	if err := c.Append(*new(float32)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(complex64)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestFloat32ColLen(t *testing.T) {
	c := new(Float32Col)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(float32))
	c.Append(*new(float32))
	c.Append(*new(float32))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}
