package circbuf

// Circbuf is a simple circular buffer with a predefined capacity. It works
// like a slice, but once full it wraps.  Operations on this object are not
// thread-safe.
type Circbuf struct {
	items []interface{}
	len   int
	cap   int
	head  int
}

// New returns an empty circular buffer with the given capacity.
func New(cap int) *Circbuf {
	if cap < 1 {
		panic("runtime error: circbuf.New: len out of range")
	}
	return &Circbuf{len: 0, cap: cap, items: make([]interface{}, cap)}
}

// Add appends an item to the buffer. If the buffer is full (length ==
// capacity), the first item in the buffer will be ejected to make room.
func (c *Circbuf) Add(e interface{}) {
	c.items[(c.head+c.len)%c.cap] = e
	if c.len < c.cap {
		c.len++
	} else {
		c.head = (c.head + 1) % c.cap
	}
}

// Do calls function f on each item of the ring, in forward order. The
// behavior of Do is undefined if f changes *r.
func (c *Circbuf) Do(f func(interface{})) {
	for i := 0; i < c.len; i++ {
		f(c.items[(c.head+i)%c.cap])
	}
}

// Len returns the number of items.
func (c *Circbuf) Len() int {
	return c.len
}

// Cap returns the number of slots.
func (c *Circbuf) Cap() int {
	return c.cap
}
