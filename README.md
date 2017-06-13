# Circ Buf

`circbuf` implements a simple circular buffer. Operations in this library are not thread-safe.

### Installation

    go get github.com/gmcnaughton/circbuf

### Usage

```go
  import (
    "circbuf"
    "fmt"
  )

  buf := circbuf.New(2)
  buf.Length()   // => 0
  buf.Capacity() // => 2

  buf.Add(1)
  buf.Add(2)
  buf.Add(3)
  buf.ForEach(func(el interface{}) {
    fmt.Println(el) // => 2, 3
  })
```

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/gmcnaughton/circbuf.

## License

The package is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
