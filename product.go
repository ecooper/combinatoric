package combinatoric

import (
	"math/big"
)

type ProductIterator struct {
	pools   [][]interface{}
	n       int
	indices []int

	max   uint64
	iters uint64

	res []interface{}
}

func (iter *ProductIterator) Next() []interface{} {
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

func (iter *ProductIterator) HasNext() bool {
	return iter.iters < iter.max
}

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

	return LenProduct(sizes...)
}

func (iter *ProductIterator) Len() uint64 {
	return iter.max
}

func LenProduct(pools ...int) uint64 {
	t := big.NewInt(1)
	for i := range pools {
		t.Mul(t, big.NewInt(int64(pools[i])))
	}
	return t.Uint64()
}

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
