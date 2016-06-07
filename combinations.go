package combinatoric

import (
	"math/big"
)

type CombinationIterator struct {
	pool    []interface{}
	n       int
	r       int
	indices []int

	max   uint64
	iters uint64

	res []interface{}
}

func (iter *CombinationIterator) Next() []interface{} {
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

func (iter *CombinationIterator) HasNext() bool {
	return iter.iters < iter.max
}

func (iter *CombinationIterator) Reset() {
	iter.iters = 0

	iter.indices = make([]int, iter.r)
	for i := 0; i < iter.r; i++ {
		iter.indices[i] = i
	}

}

func (iter *CombinationIterator) Len() uint64 {
	return iter.max
}

func LenCombinations(n int, r int) uint64 {
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

func Combinations(pool []interface{}, r int) (*CombinationIterator, error) {
	n := len(pool)

	if r > n {
		return nil, IteratorResultSizeError
	}

	iter := &CombinationIterator{
		pool: pool,
		n:    n,
		r:    r,
		res:  make([]interface{}, r, r),
		max:  LenCombinations(n, r),
	}

	iter.Reset()

	return iter, nil
}

// Type casting to insure CombinationIterator implements Iterator.
var _ Iterator = (*CombinationIterator)(nil)
