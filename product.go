package combinatoric

import (
	"math/big"
)

type ProductIterator struct {
	pools   [][]interface{}
	n       int
	indices []int

	max   *big.Int
	iters *big.Int

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

	iter.iters.Add(iter.iters, bigIntIncr)

	return iter.res
}

func (iter *ProductIterator) HasNext() bool {
	return iter.iters.Cmp(iter.max) == -1
}

func (iter *ProductIterator) Reset() {
	iter.iters = big.NewInt(0)

	iter.indices = make([]int, iter.n, iter.n)
	iter.res = make([]interface{}, iter.n, iter.n)
}

func (iter *ProductIterator) len() *big.Int {
	sizes := make([]int, iter.n, iter.n)
	for i := range sizes {
		sizes[i] = len(iter.pools[i])
	}

	return LenProduct(sizes...)
}

func (iter *ProductIterator) Len() *big.Int {
	return iter.max
}

func LenProduct(pools ...int) *big.Int {
	t := big.NewInt(1)
	for i := range pools {
		t.Mul(t, big.NewInt(int64(pools[i])))
	}
	return t
}

func Product(pools [][]interface{}) *ProductIterator {
	iter := &ProductIterator{
		pools: pools,
		n:     len(pools),
		iters: big.NewInt(0),
	}

	iter.max = iter.len()

	iter.Reset()

	return iter
}
