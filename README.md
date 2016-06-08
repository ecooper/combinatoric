# Combinatoric [![Build Status](https://travis-ci.org/ecooper/combinatoric.svg?branch=master)](https://travis-ci.org/ecooper/combinatoric) [![Coverage Status](https://coveralls.io/repos/github/ecooper/combinatoric/badge.svg?branch=master)](https://coveralls.io/github/ecooper/combinatoric?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/ecooper/combinatoric?r=1)](https://goreportcard.com/report/github.com/ecooper/combinatoric)

Combinatoric is a simple Go port of the "combinatoric" parts of Python's
[`itertools`](https://docs.python.org/3.3/library/itertools.html)--specifically, [`combinations`](https://docs.python.org/3.3/library/itertools.html#itertools.combinations), [`permutations`](https://docs.python.org/3.3/library/itertools.html#itertools.permutations), and [`product`](https://docs.python.org/3.3/library/itertools.html#itertools.product).

See [godocs](https://godoc.org/github.com/ecooper/combinatoric) for
more.


## Installation

```bash
$ go get https://github.com/ecooper/combinatoric
```

## Quickstart

```go
package main

import (
    "fmt"
    "github.com/ecooper/combinatoric"
)

func main() {
    src := []interface{}{"A", "B", "C", "D"}

    // Create a new CombinationIterator of 2 elements using src
    iter, _ := combinatoric.Combinations(src, 2)

    // Print the length of the iterator
    fmt.Printf("Expecting %d combinations:\n", iter.Len())

    // Set c to the next combination until Next returns a nil slice.
    for c := iter.First(); c != nil; c = iter.Next() {
        fmt.Printf("%s\n", c)
    }

    // Restore the CombinationsIterator to its original state
    iter.Reset()
}
```

## Usage

See [godocs](https://godoc.org/github.com/ecooper/combinatoric) for more
documentation and usage examples.
