// Test case for Complex64Col
// This file was generated. See generator directory.
package column

import "testing"

func TestComplex64ColAppend(t *testing.T) {
	c := new(Complex64Col)
	if err := c.Append(*new(complex64)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(complex128)); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func TestComplex64ColLen(t *testing.T) {
	c := new(Complex64Col)
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new(complex64))
	c.Append(*new(complex64))
	c.Append(*new(complex64))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}
