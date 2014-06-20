// Test case for {{.CapsName}}
// This file was generated. See generator directory.
package column

import "testing"

func Test{{.CapsName}}Append(t *testing.T) {
	c := new({{.CapsName}})
	if err := c.Append(*new({{.Type}})); err != nil {
		t.Fatal(err)
	}
	if err := c.Append(*new({{.AnotherType}})); err == nil {
		t.Fatal("Was expecting an error")
	}
}

func Test{{.CapsName}}Len(t *testing.T) {
	c := new({{.CapsName}})
	if c.Len() != 0 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 0)
	}
	c.Append(*new({{.Type}}))
	c.Append(*new({{.Type}}))
	c.Append(*new({{.Type}}))
	if c.Len() != 3 {
		t.Fatalf("Wrong length. (actual) %d != %d (expected)", c.Len(), 3)
	}
}
