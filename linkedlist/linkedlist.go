package linkedlist

// LinkedList is a struct that represents a linked list of items.
type LinkedList struct {
	head   *linkedListItem
	length int
}

type linkedListItem struct {
	val  interface{}
	next *linkedListItem
}

// New returns an empty list.
func New() *LinkedList {
	return &LinkedList{}
}

// Append appends the value to the end of the list. Performance is constant time.
func (l *LinkedList) Append(val interface{}) {
	item := &linkedListItem{val: val}

	tail := l.tail()
	if tail != nil {
		tail.next = item
	} else {
		l.head = item
	}
	l.length++
}

// Remove removes the value at the given index.
func (l *LinkedList) Remove(index int) interface{} {
	prev, item := l.at(index)
	if prev != nil {
		prev.next = item.next
	} else {
		l.head = item.next
	}
	l.length--
	return item.val
}

// Length returns the length of the list.
func (l *LinkedList) Length() int {
	return l.length
}

// At returns the value at index N. Performance is linear time.
func (l *LinkedList) At(index int) interface{} {
	_, item := l.at(index)
	return item.val
}

// Do calls function f on each item in the list, in forward order. The
// behavior of Do is undefined if f changes *l.
func (l *LinkedList) Do(f func(val interface{})) {
	item := l.head
	for item != nil {
		if f != nil {
			f(item.val)
		}
		item = item.next
	}
}

func (l *LinkedList) at(index int) (prev *linkedListItem, item *linkedListItem) {
	if index < 0 || index >= l.length {
		panic("runtime error: linkedlist.At: invalid index")
	}
	item = l.head
	for ; index > 0; index-- {
		prev = item
		item = item.next
	}
	return
}

// If we kept a tail pointer this would be unnecessary - and avoids making Append linear-time
func (l *LinkedList) tail() *linkedListItem {
	item := l.head
	if item == nil {
		return nil
	}

	for item.next != nil {
		item = item.next
	}
	return item
}
