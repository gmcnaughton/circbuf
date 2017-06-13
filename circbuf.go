package circbuf

type Circbuf struct {
	elements []interface{}
	len      uint
	cap      uint
	head     uint
}

// New returns a new circular buffer with the given allocated capacity.
// TODO: specify both len and cap as params?
// TODO: optional parameters
func New(cap uint) *Circbuf {
	return &Circbuf{len: 0, cap: cap, elements: make([]interface{}, cap, cap)}
}

// Add adds an element to the buffer. If the buffer is full (length == capacity),
// Add removes the first item and appends the new item at the end.
func (c *Circbuf) Add(e interface{}) {
	c.elements[(c.head+c.len)%c.cap] = e
	if c.len < c.cap {
		c.len++
	} else {
		c.head = (c.head + 1) % c.cap
	}
}

// ForEach calls the given function once for each element in the buffer.
// TODO: allow function to return an error to halt?
func (c *Circbuf) ForEach(f func(interface{})) {
	for i := uint(0); i < c.len; i++ {
		f(c.elements[(c.head+i)%c.cap])
	}
}

// Length returns the number of elements which are currently in the buffer.
// It will always be <= the buffer's capacity.
func (c *Circbuf) Length() uint {
	return c.len
}

// Capacity returns the number of elements which can fit in the buffer. It will
// always be >= the buffer's length.
func (c *Circbuf) Capacity() uint {
	return c.cap
}
