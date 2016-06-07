package combinatoric

import (
	"math/big"
)

// ProductIterator implements an Iterator to generate the Cartesian
// product of a slice of slices.
//
// ProductIterator is not threadsafe, and should always be initialized
// through Products. Return values from First and Next share the same
// memory address. Use copy if the value must persist between
// iterations.
type ProductIterator struct {
	pools   [][]interface{}
	n       int
	indices []int

	max   uint64
	iters uint64

	res []interface{}
}

// First resets the iterator to its starting state and returns the first
// product.
func (iter *ProductIterator) First() []interface{} {
	iter.Reset()
	return iter.Next()
}

// Next returns the next value in the iterator or returns nil.
func (iter *ProductIterator) Next() []interface{} {
	if !iter.HasNext() {
		return nil
	}

	if iter.res[0] == nil {
		for i := range iter.res {
			iter.res[i] = iter.pools[i][0]
		}
	} else {
		for i := iter.n - 1; i >= 0; i-- {
			pool := iter.pools[i]
			iter.indices[i] += 1
			if iter.indices[i] == len(pool) {
				iter.indices[i] = 0
				iter.res[i] = pool[0]
			} else {
				iter.res[i] = pool[iter.indices[i]]
				break
			}
		}
	}

	iter.iters += 1

	return iter.res
}

// HasNext returns true if the iterator is not yet exhausted.
func (iter *ProductIterator) HasNext() bool {
	return iter.iters < iter.max
}

// Reset returns the iterator to its starting state.
func (iter *ProductIterator) Reset() {
	iter.iters = 0

	iter.indices = make([]int, iter.n, iter.n)
	iter.res = make([]interface{}, iter.n, iter.n)
}

func (iter *ProductIterator) len() uint64 {
	sizes := make([]int, iter.n, iter.n)
	for i := range sizes {
		sizes[i] = len(iter.pools[i])
	}

	return lenProduct(sizes...)
}

// Len returns the maximum iterations.
func (iter *ProductIterator) Len() uint64 {
	return iter.max
}

func lenProduct(pools ...int) uint64 {
	t := big.NewInt(1)
	for i := range pools {
		t.Mul(t, big.NewInt(int64(pools[i])))
	}
	return t.Uint64()
}

// Product creates a new ProductIterator for a given slice and
// desired output size.
func Product(pools [][]interface{}) (*ProductIterator, error) {
	iter := &ProductIterator{
		pools: pools,
		n:     len(pools),
	}

	iter.max = iter.len()

	iter.Reset()

	return iter, nil
}

// Type casting to insure ProductIterator implements Iterator.
var _ Iterator = (*ProductIterator)(nil)
