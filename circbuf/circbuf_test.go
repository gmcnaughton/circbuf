package circbuf

import (
	"errors"
	"testing"
)

func TestNewWithNegativeSize(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected New with negative size to panic")
		}
	}()

	New(-1)
	t.Error("Expected New with negative size to panic")
}

func TestNewWithZeroSize(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected New with size 0 to panic")
		}
	}()

	New(0)
	t.Error("Expected New with size 0 to panic")
}

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
	AssertEqual(t, c, []interface{}{1, 2})
}

func TestAddOverCapacity(t *testing.T) {
	c := New(2)
	c.Add(1)
	c.Add(2)
	c.Add(3)
	AssertEqual(t, c, []interface{}{2, 3})
	if c.Len() != c.Cap() {
		t.Error("Expected wrapped buffer to have length == capacity", c.Len(), c.Cap())
	}

	c.Add(4)
	AssertEqual(t, c, []interface{}{3, 4})
	if c.Len() != c.Cap() {
		t.Error("Expected wrapped buffer to have length == capacity", c.Len(), c.Cap())
	}
}

func TestDoWithEmptyBuffer(t *testing.T) {
	c := New(2)
	err := c.Do(func(item interface{}) error {
		t.Error("Expected Do on empty buffer never to get called")
		return nil
	})
	if err != nil {
		t.Error("Expected Do on empty buffer to not return an error", err)
	}
}

func TestDoWithBufferWithOneItem(t *testing.T) {
	c := New(2)
	c.Add(1)
	count := 0
	err := c.Do(func(item interface{}) error {
		if item != 1 {
			t.Error("Expected Do to get called with items in buffer", 1, item)
		}
		count++
		return nil
	})
	if err != nil {
		t.Error("Expected Do on empty buffer to not return an error", err)
	}
	if count != c.Len() {
		t.Error("Expected Do to get called once per item in buffer")
	}
}

func TestDoHaltsWhenFuncReturnsError(t *testing.T) {
	c := New(2)
	c.Add(1)
	c.Add(2)
	sampleerr := errors.New("sample")
	count := 0
	err := c.Do(func(item interface{}) error {
		if count > 0 {
			return sampleerr
		}
		count++
		return nil
	})
	if count != 1 {
		t.Error("Expected Do to halt when error is returned")
	}
	if err != sampleerr {
		t.Error("Expected Do to return the error from the wrapper func after halting")
	}
}

func TestSliceWithEmptyBuffer(t *testing.T) {
	c := New(2)
	expected, actual := []interface{}{}, c.Slice()
	AssertEqualSlice(t, actual, expected)
}

func TestSliceWithBufferWithOneItem(t *testing.T) {
	c := New(2)
	c.Add(1)
	expected, actual := []interface{}{1}, c.Slice()
	AssertEqualSlice(t, actual, expected)
}

func AssertEqual(t *testing.T, c *Circbuf, expected []interface{}) {
	actual := c.Slice()
	AssertEqualSlice(t, actual, expected)
}

func AssertEqualSlice(t *testing.T, actual []interface{}, expected []interface{}) {
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
