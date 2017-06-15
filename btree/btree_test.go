package btree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testItem int

func (c testItem) Compare(o Item) int {
	other, ok := o.(testItem)
	if !ok {
		panic("can't compare int to non-int")
	}
	if c > other {
		return 1
	} else if c == other {
		return 0
	} else { // c < other
		return -1
	}
}

func empty() *Btree {
	return New()
}

// 1
//  1
//   1
func linear() *Btree {
	b := New()
	b.Add(testItem(1))
	b.Add(testItem(1))
	b.Add(testItem(1))
	return b
}

// 1
//  2
//   3
func rightHanded() *Btree {
	b := New()
	b.Add(testItem(1))
	b.Add(testItem(2))
	b.Add(testItem(3))
	return b
}

//   2
//  1 3
func balanced() *Btree {
	b := New()
	b.Add(testItem(2))
	b.Add(testItem(3))
	b.Add(testItem(1))
	return b
}

//   1
//  2
// 3
func leftHanded() *Btree {
	b := New()
	b.Add(testItem(3))
	b.Add(testItem(2))
	b.Add(testItem(1))
	return b
}

//                100
//            50      150
//          49     149  150
//        48              150
//                          151
//                        150
func complex() *Btree {
	b := New()
	b.Add(testItem(100))
	b.Add(testItem(50))
	b.Add(testItem(150))
	b.Add(testItem(149))
	b.Add(testItem(150))
	b.Add(testItem(150))
	b.Add(testItem(151))
	b.Add(testItem(150))
	b.Add(testItem(49))
	b.Add(testItem(48))
	return b
}

//    1
// *2
func invalidLeft() *Btree {
	b := New()
	b.head = &node{val: testItem(1)}
	b.head.left = &node{val: testItem(2)}
	return b
}

//  1
//   *0
func invalidRight() *Btree {
	b := New()
	b.head = &node{val: testItem(1)}
	b.head.right = &node{val: testItem(0)}
	return b
}

//          100
//     50         150
//       75   *74
func invalidComplex() *Btree {
	b := New()
	b.head = &node{val: testItem(100)}
	b.head.left = &node{val: testItem(50)}
	b.head.left.right = &node{val: testItem(75)}
	b.head.right = &node{val: testItem(150)}
	b.head.right.left = &node{val: testItem(74)}
	return b
}

func TestNewIsEmpty(t *testing.T) {
	b := New()
	assert.Equal(t, 0, b.Size(), "New should be empty")
}

func TestAddIncreasesSize(t *testing.T) {
	b := New()
	b.Add(testItem(1))
	b.Add(testItem(2))
	assert.Equal(t, 2, b.Size(), "Add should increase size")
}

func TestAdd(t *testing.T) {
	tests := []struct {
		b        *Btree
		expected string
	}{
		{empty(), "()"},
		{linear(), "(1, (1, (1)))"},
		{rightHanded(), "(1, (2, (3)))"},
		{balanced(), "((1), 2, (3))"},
		{leftHanded(), "(((1), 2), 3)"},
		{complex(), "((((48), 49), 50), 100, ((149), 150, (150, (150, ((150), 151)))))"},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.b.String(), "String did not work")
	}
}

func TestInclude(t *testing.T) {
	tests := []struct {
		b          *Btree
		expected   []testItem
		unexpected []testItem
	}{
		{empty(), []testItem{}, []testItem{1, 2, 3}},
		{linear(), []testItem{1}, []testItem{4, -1}},
		{rightHanded(), []testItem{1, 2, 3}, []testItem{4, -1}},
		{balanced(), []testItem{1, 2, 3}, []testItem{4, -1}},
		{leftHanded(), []testItem{1, 2, 3}, []testItem{4, -1}},
		{complex(), []testItem{48, 49, 50, 100, 149, 150, 151}, []testItem{4, -1}},
	}
	for _, tt := range tests {
		for _, val := range tt.expected {
			assert.Equal(t, true, tt.b.Include(val), "Expected value not included")
		}
		for _, val := range tt.unexpected {
			assert.Equal(t, false, tt.b.Include(val), "Unexpected val included")
		}
	}
}

func TestDepthFirst(t *testing.T) {
	tests := []struct {
		b        *Btree
		expected []testItem
	}{
		{empty(), []testItem{}},
		{linear(), []testItem{1, 1, 1}},
		{rightHanded(), []testItem{1, 2, 3}},
		{balanced(), []testItem{1, 2, 3}},
		{leftHanded(), []testItem{1, 2, 3}},
		{complex(), []testItem{48, 49, 50, 100, 149, 150, 150, 150, 150, 151}},
	}
	for _, tt := range tests {
		count := 0
		tt.b.DepthFirst(func(item Item) {
			assert.Equal(t, tt.expected[count], item, "Depth first visits nodes in DFS order")
			count++
		})
		assert.Equal(t, len(tt.expected), count, "Calls f once per item in tree")
	}
}

func TestBreadthFirst(t *testing.T) {
	tests := []struct {
		b        *Btree
		expected []testItem
	}{
		{empty(), []testItem{}},
		{linear(), []testItem{1, 1, 1}},
		{rightHanded(), []testItem{1, 2, 3}},
		{balanced(), []testItem{2, 1, 3}},
		{leftHanded(), []testItem{3, 2, 1}},
		{complex(), []testItem{100, 50, 150, 49, 149, 150, 48, 150, 151, 150}},
	}
	for _, tt := range tests {
		count := 0
		tt.b.BreadthFirst(func(item Item) {
			assert.Equal(t, tt.expected[count], item, "Breadth first visits nodes in BFS order")
			count++
		})
		assert.Equal(t, len(tt.expected), count, "Calls f once per item in tree")
	}
}

func TestValid(t *testing.T) {
	tests := []struct {
		b        *Btree
		expected bool
	}{
		{empty(), true},
		{linear(), true},
		{rightHanded(), true},
		{balanced(), true},
		{leftHanded(), true},
		{complex(), true},
		{invalidLeft(), false},
		{invalidRight(), false},
		{invalidComplex(), false},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.expected, tt.b.Valid(), fmt.Sprintf("Valid did not work: %v", tt.b))
	}
}
