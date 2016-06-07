package combinatoric

import (
	"errors"
)

type Iterator interface {
	First() []interface{}
	Next() []interface{}
	HasNext() bool
	Len() uint64
	Reset()
}

var IteratorResultSizeError error = errors.New("Result size is larger than input")
