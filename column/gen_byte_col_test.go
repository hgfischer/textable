// Test case for ByteCol
// This file was generated. See generator directory.
package column

import (
	"fmt"
	"testing"
)

func TestByteColAppend(t *testing.T) {
	c := new(ByteCol)
	if err := c.Append(*new(byte)); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new(uint)); err == nil {
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

func TestByteColAt(t *testing.T) {
	c := new(ByteCol)
	expected := *new(byte)
	c.Append(expected)
	value, exists := c.At(0)
	if !exists {
		t.Fatal("Value at index 0 should exist")
	}
	if value != expected {
		t.Fatalf("Wrong value. (actual) %#v != %#v (expected)", value, expected)
	}
	if value, exists = c.At(1); exists {
		t.Fatal("Value at index 1 should not exist")
	}
}

func TestByteColString(t *testing.T) {
	c := new(ByteCol)
	val := *new(byte)
	c.Append(val)
	c.Append(val)
	expected := fmt.Sprintf("[]byte{%#v, %#v}", val, val)
	value := c.String()
	if value != expected {
		t.Fatalf("Wrong value. (actual) %#v != %#v (expected)", value, expected)
	}
}
