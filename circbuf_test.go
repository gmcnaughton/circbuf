package circbuf

import (
	"testing"
)

func TestLengthOfEmptyBuffer(t *testing.T) {
	c := New(1)
	if c.Length() != 0 {
		t.Error("Expected empty buffer to have length 0, was", c.Length())
	}
}

func TestLengthOfBufferWithOneElement(t *testing.T) {
	c := New(1)
	c.Add(1)
	if c.Length() != 1 {
		t.Error("Expected buffer with one element to have length 1, was", c.Length())
	}
}

func TestCapacityOfEmptyBuffer(t *testing.T) {
	c := New(1)
	if c.Capacity() != 1 {
		t.Error("Expected empty buffer to have capacity 1, was", c.Capacity())
	}
}

func TestCapacityOfBufferWithOneElement(t *testing.T) {
	c := New(1)
	if c.Capacity() != 1 {
		t.Error("Expected buffer with one element to have capacity 1, was", c.Capacity())
	}
}

func TestAdd(t *testing.T) {
	c := New(2)
	c.Add(1)
	c.Add(2)
	AssertEqualInt(t, c, []int{1, 2})
}

func TestAddWraps(t *testing.T) {
	c := New(2)
	c.Add(1)
	c.Add(2)
	c.Add(3)
	AssertEqualInt(t, c, []int{2, 3})
	if c.Length() != c.Capacity() {
		t.Error("Expected wrapped buffer to have length == capacity", c.Length(), c.Capacity())
	}

	c.Add(4)
	AssertEqualInt(t, c, []int{3, 4})
	if c.Length() != c.Capacity() {
		t.Error("Expected wrapped buffer to have length == capacity", c.Length(), c.Capacity())
	}
}

func AssertEqualInt(t *testing.T, c *Circbuf, expected []int) {
	actual := make([]interface{}, 0, c.Length())
	c.ForEach(func(el interface{}) {
		actual = append(actual, el)
	})

	if len(actual) != len(expected) {
		t.Errorf("Expected buffer contents (%v) to equal (%v)", actual, expected)
		return
	}

	for i, el := range actual {
		if el != expected[i] {
			t.Errorf("Expected buffer contents (%v) to equal (%v)", actual, expected)
			return
		}
	}
}
