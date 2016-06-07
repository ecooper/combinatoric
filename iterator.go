package combinatoric

import (
	"errors"
)

type Iterator interface {
	HasNext() bool
	Next() []interface{}
	Len() uint64
	Reset()
}

var IteratorResultSizeError error = errors.New("Result size is larger than input")
