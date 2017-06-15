package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewIsEmpty(t *testing.T) {
	l := New()
	assert.Equal(t, 0, l.Length(), "New list should be empty")
}

func TestAppendMakesLonger(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	assert.Equal(t, 3, l.Length(), "Appending to list should make it longer")
}

func TestAppendAddsAtEnd(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	assert.Equal(t, 1, l.At(0), "Appending to list should add value at end")
	assert.Equal(t, 2, l.At(1), "Appending to list should add value at end")
	assert.Equal(t, 3, l.At(2), "Appending to list should add value at end")
}

func TestAtWithNegativeIndex(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected At with negative index to panic")
		}
	}()

	l := New()
	l.At(-1)
	t.Error("Expected At with negative index to panic")
}

func TestAtWithIndexAtEnd(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected At with index at end to panic")
		}
	}()

	l := New()
	l.Append(1)
	l.At(1)
	t.Error("Expected At with index at end to panic")
}

func TestAtWithIndexAfterEnd(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected At with index past end to panic")
		}
	}()

	l := New()
	l.Append(1)
	l.At(100)
	t.Error("Expected At with index past end to panic")
}

func TestRemoveMakesShorter(t *testing.T) {
	l := New()
	l.Append(1)
	l.Remove(0)
	assert.Equal(t, 0, l.Length(), "Removing from a list should make it shorter")
}

func TestRemoveRemovesValue(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Remove(1)
	assert.Equal(t, 1, l.At(0), "Removing from a list should remove the requested item")
	assert.Equal(t, 3, l.At(1), "Removing from a list should remove the requested item")
}

func TestRemoveReturnsValueAtIndex(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	val := l.Remove(1)
	assert.Equal(t, 2, val, "Removing from a list should return the value removed")
}

func TestRemoveWithNegativeIndex(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected Remove with negative index to panic")
		}
	}()

	l := New()
	l.Remove(-1)
	t.Error("Expected Remove with negative index to panic")
}

func TestRemoveWithIndexAtEnd(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected Remove with index at end to panic")
		}
	}()

	l := New()
	l.Append(1)
	l.Remove(1)
	t.Error("Expected Remove with index at end to panic")
}

func TestRemoveWithIndexAfterEnd(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected Remove with index past end to panic")
		}
	}()

	l := New()
	l.Append(1)
	l.Remove(100)
	t.Error("Expected Remove with index past end to panic")
}

func TestDoWithEmptyList(t *testing.T) {
	l := New()
	count := 0
	l.Do(func(val interface{}) {
		count++
	})
	assert.Equal(t, 0, count, "Do on empty list should never call func")
}

func TestDo(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append("foo")
	count := 0
	l.Do(func(val interface{}) {
		switch count {
		case 0:
			assert.Equal(t, 1, val, "Do on empty list should call with each val in forward order")
		case 1:
			assert.Equal(t, 2, val, "Do on empty list should call with each val in forward order")
		case 2:
			assert.Equal(t, "foo", val, "Do on empty list should call with each val in forward order")
		}
		count++
	})
	assert.Equal(t, 3, count, "Do on list should be called once per value")
}
