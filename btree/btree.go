package btree

import (
	"fmt"

	"github.com/golang-collections/go-datastructures/queue"
)

// Btree implements a binary search tree.
type Btree struct {
	head *node
	size int
}

type node struct {
	left  *node
	right *node
	val   Item
}

// Visitor defines a func which can be called on any item in the tree.
type Visitor func(val Item)

// Item defines an interface values which can be added to a Btree.
type Item interface {
	// Compare returns a bool that can be used to determine
	// ordering in the priority queue.  Assuming the queue
	// is in ascending order, this should return > logic.
	// Return 1 to indicate this object is greater than the
	// the other logic, 0 to indicate equality, and -1 to indicate
	// less than other.
	Compare(other Item) int
}

// New returns an empty btree
func New() *Btree {
	return &Btree{}
}

// Add inserts the given value in the btree
func (b *Btree) Add(val Item) {
	new := &node{val: val}

	if b.head == nil {
		b.head = new
	} else {
		b.head.add(new)
	}
	b.size++
}

func (n *node) add(new *node) {
	cmp := new.val.Compare(n.val)
	if cmp >= 0 {
		// new value is equal to or larger than current value, put on the right
		if n.right == nil {
			n.right = new
		} else {
			n.right.add(new)
		}
	} else {
		// new value is less than the current value, put on the left
		if n.left == nil {
			n.left = new
		} else {
			n.left.add(new)
		}
	}
}

func (b *Btree) String() string {
	if b.head == nil {
		return "()"
	}
	return b.head.String()
}

func (n *node) String() string {
	if n.left != nil && n.right != nil {
		return fmt.Sprintf("(%v, %v, %v)", n.left, n.val, n.right)
	} else if n.left != nil {
		return fmt.Sprintf("(%v, %v)", n.left, n.val)
	} else if n.right != nil {
		return fmt.Sprintf("(%v, %v)", n.val, n.right)
	}
	return fmt.Sprintf("(%v)", n.val)
}

// Size returns the number of values in the btree
func (b *Btree) Size() int {
	return b.size
}

// Include returns true if the btree includes the value and false otherwise
func (b *Btree) Include(val Item) bool {
	if b.head == nil {
		return false
	}

	return b.head.include(val)
}

func (n *node) include(val Item) bool {
	cmp := val.Compare(n.val)
	if cmp == 0 {
		return true
	} else if cmp < 0 {
		if n.left == nil {
			return false
		}
		return n.left.include(val)
	} else {
		if n.right == nil {
			return false
		}
		return n.right.include(val)
	}
}

// DepthFirst calls function f on each item in the tree, in depth-first order.
// The behavior is undefined if f changes *b.
func (b *Btree) DepthFirst(f Visitor) {
	b.head.depthFirst(f)
}

func (n *node) depthFirst(f Visitor) {
	if n == nil {
		return
	}

	n.left.depthFirst(f)
	if f != nil {
		f(n.val)
	}
	n.right.depthFirst(f)
}

// BreadthFirst calls function f on each item in the tree, in breadth-first order.
// The behavior is undefined if f changes *b.
func (b *Btree) BreadthFirst(f Visitor) {
	q := queue.New(int64(b.Size()))

	q.Put(b.head)
	for !q.Empty() {
		ns, _ := q.Get(1)
		n := ns[0].(*node)
		if n != nil {
			if f != nil {
				f(n.val)
			}
			q.Put(n.left)
			q.Put(n.right)
		}
	}
}

// Valid returns true if this is a valid binary tree (sorted).
func (b *Btree) Valid() bool {
	ok, _ := b.head.valid(nil)
	return ok
}

func (n *node) valid(max Item) (bool, Item) {
	if n == nil {
		return true, max
	}

	ok, max := n.left.valid(max)
	if !ok {
		return false, max
	}

	if max != nil && max.Compare(n.val) > 0 {
		return false, max
	}
	max = n.val

	ok, max = n.right.valid(max)
	return ok, max
}
