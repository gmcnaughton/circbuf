package circbuf

// Circbuf is a simple circular buffer with a predefined capacity. It works
// like a slice, but wraps when full. Operations on this object are not
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
// capacity), the first item is removed to make room.
func (c *Circbuf) Add(e interface{}) {
	c.items[(c.head+c.len)%c.cap] = e
	if c.len < c.cap {
		c.len++
	} else {
		c.head = (c.head + 1) % c.cap
	}
}

// Do calls function f on each item of the buffer, in forward order. The
// behavior of Do is undefined if f changes *c.
func (c *Circbuf) Do(f func(interface{})) {
	for i := 0; i < c.len; i++ {
		f(c.items[(c.head+i)%c.cap])
	}
}

// Len returns the number of items in the buffer.
func (c *Circbuf) Len() int {
	return c.len
}

// Cap returns the number of items the buffer can hold before wrapping.
func (c *Circbuf) Cap() int {
	return c.cap
}

// Slice returns each item of the buffer, in forward order. Runs in linear time,
// as all items must be copied to a temporary slice in order to linearize the
// buffer.
func (c *Circbuf) Slice() []interface{} {
	items := make([]interface{}, 0, c.Len())
	c.Do(func(item interface{}) {
		items = append(items, item)
	})
	return items
}
