package combinatoric

import (
	"errors"
	"math/big"
)

// CombinationIterator implements an Iterator to generate unique
// combinations of a given slice.
//
// CombinationIterator is not threadsafe, and should always be
// initialized through Combinations. Return values from First and Next
// share the same memory address. Use copy if the value must persist
// between iterations.
type CombinationIterator struct {
	pool    []interface{}
	n       int
	r       int
	indices []int

	max   uint64
	iters uint64

	res []interface{}
}

// First resets the iterator to its starting state and returns the first
// combination.
func (iter *CombinationIterator) First() []interface{} {
	iter.Reset()
	return iter.Next()
}

// Next returns the next value in the iterator or returns nil.
func (iter *CombinationIterator) Next() []interface{} {
	if !iter.HasNext() {
		return nil
	}

	if iter.res[0] != nil {
		i := iter.r - 1
		for ; i > -1; i-- {
			if iter.indices[i] != i+iter.n-iter.r {
				break
			}
		}
		if i > -1 {
			iter.indices[i] += 1
			for j := i + 1; j < iter.r; j++ {
				iter.indices[j] = iter.indices[j-1] + 1
			}
		}
	}

	for j := 0; j < iter.r; j++ {
		iter.res[j] = iter.pool[iter.indices[j]]
	}

	iter.iters += 1

	return iter.res
}

// HasNext returns true if the iterator is not yet exhausted.
func (iter *CombinationIterator) HasNext() bool {
	return iter.iters < iter.max
}

// Reset returns the iterator to its starting state.
func (iter *CombinationIterator) Reset() {
	iter.iters = 0

	iter.indices = make([]int, iter.r)
	for i := 0; i < iter.r; i++ {
		iter.indices[i] = i
	}

	iter.res = make([]interface{}, iter.r, iter.r)
}

// Len returns the maximum iterations.
func (iter *CombinationIterator) Len() uint64 {
	return iter.max
}

func lenCombinations(n int, r int) uint64 {
	n64 := int64(n)
	r64 := int64(r)

	if n < r {
		return 0
	}

	total := big.NewInt(0)

	total.Mul(
		factorial(big.NewInt(n64-r64)),
		factorial(big.NewInt(r64)),
	)

	total.Div(factorial(big.NewInt(n64)), total)

	return total.Uint64()
}

// Combinations creates a new CombinationIterator for a given slice and
// desired output size.
func Combinations(pool []interface{}, r int) (*CombinationIterator, error) {
	n := len(pool)

	if r > n {
		return nil, errors.New("Result size is larger than input")
	}

	iter := &CombinationIterator{
		pool: pool,
		n:    n,
		r:    r,
		res:  make([]interface{}, r, r),
		max:  lenCombinations(n, r),
	}

	iter.Reset()

	return iter, nil
}

// Type casting to insure CombinationIterator implements Iterator.
var _ Iterator = (*CombinationIterator)(nil)
