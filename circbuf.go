package circbuf

// Circbuf is a simple circular buffer with a predefined capacity. It works
// like a slice, but once full it wraps.  Operations on this object are not
// thread-safe.
type Circbuf struct {
	elements []interface{}
	len      uint
	cap      uint
	head     uint
}

// New returns an empty circular buffer with the given capacity.
func New(cap uint) *Circbuf {
	// TODO: specify both len and cap as params?
	// TODO: optional parameters?
	return &Circbuf{len: 0, cap: cap, elements: make([]interface{}, cap, cap)}
}

// Add appends an element to the buffer. If the buffer is full (length ==
// capacity), the first item in the buffer will be ejected to make room.
func (c *Circbuf) Add(e interface{}) {
	// TODO: name 'append' or 'push' instead of 'add'?
	c.elements[(c.head+c.len)%c.cap] = e
	if c.len < c.cap {
		c.len++
	} else {
		c.head = (c.head + 1) % c.cap
	}
}

// ForEach calls the given function once for each element, in order.
// Behavior is undefined if f modifies *c.
func (c *Circbuf) ForEach(f func(interface{})) {
	// TODO: allow function to return an error to halt?
	// TODO: name 'each' instead of 'foreach'?
	for i := uint(0); i < c.len; i++ {
		f(c.elements[(c.head+i)%c.cap])
	}
}

// Length returns the number of elements.
func (c *Circbuf) Length() uint {
	return c.len
}

// Capacity returns the number of slots.
func (c *Circbuf) Capacity() uint {
	return c.cap
}
