package combinatoric

import (
	"math/big"
)

type Iterator interface {
	HasNext() bool
	Next() []interface{}
	Len() *big.Int
	Reset()
}
