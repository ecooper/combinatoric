package combinatoric

import (
	"math/big"
)

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

func (iter *PermutationIterator) Next() []interface{} {
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

func (iter *PermutationIterator) HasNext() bool {
	return iter.iters < iter.max
}

func (iter *PermutationIterator) Reset() {
	iter.iters = 0
	iter.max = iter.Len()

	iter.indices = make([]int, iter.n)
	for i := range iter.indices {
		iter.indices[i] = i
	}

	iter.cycles = make([]int, iter.n-(iter.n-iter.r))
	for i := range iter.cycles {
		iter.cycles[i] = iter.n - i
	}
}

func (iter *PermutationIterator) Len() uint64 {
	return iter.max
}

func LenPermutations(n int, r int) uint64 {
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

func Permutations(pool []interface{}, r int) (*PermutationIterator, error) {
	n := len(pool)

	if r > n {
		return nil, IteratorResultSizeError
	}

	iter := &PermutationIterator{
		pool: pool,
		n:    n,
		r:    r,
		res:  make([]interface{}, r, r),
		max:  LenPermutations(n, r),
	}

	iter.Reset()

	return iter, nil
}

// Type casting to insure PermutationIterator implements Iterator.
var _ Iterator = (*PermutationIterator)(nil)
