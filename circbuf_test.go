package circbuf

import (
	"testing"
)

func TestLenWithEmptyBuffer(t *testing.T) {
	c := New(1)
	if c.Len() != 0 {
		t.Error("Expected empty buffer to have length 0, was", c.Len())
	}
}

func TestLenWithBufferWithOneItem(t *testing.T) {
	c := New(1)
	c.Add(1)
	if c.Len() != 1 {
		t.Error("Expected buffer with one item to have length 1, was", c.Len())
	}
}

func TestCapWithEmptyBuffer(t *testing.T) {
	c := New(1)
	if c.Cap() != 1 {
		t.Error("Expected empty buffer to have capacity 1, was", c.Cap())
	}
}

func TestCapWithBufferWithOneItem(t *testing.T) {
	c := New(1)
	if c.Cap() != 1 {
		t.Error("Expected buffer with one item to have capacity 1, was", c.Cap())
	}
}

func TestAdd(t *testing.T) {
	c := New(2)
	c.Add(1)
	c.Add(2)
	AssertEqualInt(t, c, []int{1, 2})
}

func TestAddOverCapacity(t *testing.T) {
	c := New(2)
	c.Add(1)
	c.Add(2)
	c.Add(3)
	AssertEqualInt(t, c, []int{2, 3})
	if c.Len() != c.Cap() {
		t.Error("Expected wrapped buffer to have length == capacity", c.Len(), c.Cap())
	}

	c.Add(4)
	AssertEqualInt(t, c, []int{3, 4})
	if c.Len() != c.Cap() {
		t.Error("Expected wrapped buffer to have length == capacity", c.Len(), c.Cap())
	}
}

func TestDoWithEmptyBuffer(t *testing.T) {
	c := New(2)
	c.Do(func(item interface{}) {
		t.Error("Expected Do on empty buffer never to get called")
	})
}

func TestDoWithBufferWithOneItem(t *testing.T) {
	c := New(2)
	c.Add(1)
	count := uint(0)
	c.Do(func(item interface{}) {
		if item != 1 {
			t.Error("Expected Do to get called with items in buffer", 1, item)
		}
		count++
	})
	if count != c.Len() {
		t.Error("Expected Do to get called once per item in buffer")
	}
}

func AssertEqualInt(t *testing.T, c *Circbuf, expected []int) {
	actual := make([]interface{}, 0, c.Len())
	c.Do(func(item interface{}) {
		actual = append(actual, item)
	})

	if len(actual) != len(expected) {
		t.Errorf("Expected buffer contents (%v) to equal (%v)", actual, expected)
		return
	}

	for i, item := range actual {
		if item != expected[i] {
			t.Errorf("Expected buffer contents (%v) to equal (%v)", actual, expected)
			return
		}
	}
}
