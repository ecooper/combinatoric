package combinatoric

import (
	"errors"
	"math/big"
)

// PermutationIterator implements an Iterator to generate all
// permutations of a given slice.
//
// PermutationIterator is not threadsafe, and should always be
// initialized through Permutations. Return values from First and Next
// share the same memory address. Use copy if the value must persist
// between iterations.
type PermutationIterator struct {
	pool    []interface{}
	n       int
	r       int
	cycles  []int
	indices []int

	max   uint64
	iters uint64

	res []interface{}
}

// First resets the iterator to its starting state and returns the first
// permutation.
func (iter *PermutationIterator) First() []interface{} {
	iter.Reset()
	return iter.Next()
}

// Next returns the next value in the iterator or returns nil.
func (iter *PermutationIterator) Next() []interface{} {
	if !iter.HasNext() {
		return nil
	}

	if iter.res[0] != nil {
		for i := iter.r - 1; i > -1; i-- {
			iter.cycles[i] -= 1
			if iter.cycles[i] == 0 {
				for x := i; x < iter.n-1; x++ {
					v := iter.indices[x]
					iter.indices[x] = iter.indices[x+1]
					iter.indices[x+1] = v
				}
				iter.cycles[i] = iter.n - i
			} else {
				x := iter.indices[i]
				iter.indices[i] = iter.indices[iter.n-iter.cycles[i]]
				iter.indices[iter.n-iter.cycles[i]] = x
				break
			}
		}
	}

	for i := 0; i < iter.r; i++ {
		iter.res[i] = iter.pool[iter.indices[i]]
	}

	iter.iters += 1

	return iter.res
}

// HasNext returns true if the iterator is not yet exhausted.
func (iter *PermutationIterator) HasNext() bool {
	return iter.iters < iter.max
}

// Reset returns the iterator to its starting state.
func (iter *PermutationIterator) Reset() {
	iter.iters = 0

	iter.indices = make([]int, iter.n)
	for i := range iter.indices {
		iter.indices[i] = i
	}

	iter.cycles = make([]int, iter.n-(iter.n-iter.r))
	for i := range iter.cycles {
		iter.cycles[i] = iter.n - i
	}

	iter.res = make([]interface{}, iter.r, iter.r)
}

// Len returns the maximum iterations.
func (iter *PermutationIterator) Len() uint64 {
	return iter.max
}

func lenPermutations(n int, r int) uint64 {
	n64 := int64(n)
	r64 := int64(r)

	if n < r {
		return 0
	}

	total := big.NewInt(0)

	total.Div(
		factorial(big.NewInt(n64)),
		factorial(big.NewInt(n64-r64)),
	)

	return total.Uint64()
}

// Permutations creates a new PermutationIterator for a given slice and
// desired output size.
func Permutations(pool []interface{}, r int) (*PermutationIterator, error) {
	n := len(pool)

	if r > n {
		return nil, errors.New("Result size is larger than input")
	}

	iter := &PermutationIterator{
		pool: pool,
		n:    n,
		r:    r,
		res:  make([]interface{}, r, r),
		max:  lenPermutations(n, r),
	}

	iter.Reset()

	return iter, nil
}

// Type casting to insure PermutationIterator implements Iterator.
var _ Iterator = (*PermutationIterator)(nil)
