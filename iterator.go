package combinatoric

import (
	"errors"
	"math/big"
)

type Iterator interface {
	HasNext() bool
	Next() []interface{}
	Len() *big.Int
	Reset()
}

var IteratorResultSizeError error = errors.New("Result size is larger than input")
