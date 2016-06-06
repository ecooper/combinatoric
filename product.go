package combinatoric

import (
	"math/big"
)

type ProductIterator struct {
	pools      [][]interface{}
	indices    []int
	current    []interface{}
	total      int64
	iterations int64
}

func (iter *ProductIterator) Next() []interface{} {
	product := iter.EmptyProduct()
	copy(product, iter.current)

	if iter.iterations == 0 {
		for i := range product {
			product[i] = iter.pools[i][0]
		}
	} else {
		for i := len(product) - 1; i >= 0; i-- {
			pool := iter.pools[i]
			iter.indices[i] += 1
			if iter.indices[i] == len(pool) {
				iter.indices[i] = 0
				product[i] = pool[0]
			} else {
				product[i] = pool[iter.indices[i]]
				break
			}
		}
	}

	iter.current = product
	iter.iterations++
	return product
}

func (iter *ProductIterator) HasNext() bool {
	return iter.iterations < iter.total
}

func (iter *ProductIterator) EmptyProduct() []interface{} {
	return make([]interface{}, len(iter.pools))
}

func (iter *ProductIterator) Total() *big.Int {
	sizes := make([]int, len(iter.pools))
	for i := range sizes {
		sizes[i] = len(iter.pools[i])
	}

	return TotalProduct(sizes...)
}

func (iter *ProductIterator) Reset() {
	iter.iterations = 0
	iter.total = iter.Total().Int64()

	n := len(iter.pools)
	iter.indices = make([]int, n)
	iter.current = make([]interface{}, n)
}

func TotalProduct(pools ...int) *big.Int {
	t := big.NewInt(1)
	for i := range pools {
		t.Mul(t, big.NewInt(int64(pools[i])))
	}
	return t
}

func Product(pools [][]interface{}) *ProductIterator {
	p := new(ProductIterator)
	p.pools = pools
	p.Reset()
	return p
}
