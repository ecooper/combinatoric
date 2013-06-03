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

func (this *ProductIterator) Next() []interface{} {
	product := this.EmptyProduct()
	copy(product, this.current)

	if this.iterations == 0 {
		for i := range product {
			product[i] = this.pools[i][0]
		}
	} else {
		for i := len(product) - 1; i >= 0; i-- {
			pool := this.pools[i]
			this.indices[i] += 1
			if this.indices[i] == len(pool) {
				this.indices[i] = 0
				product[i] = pool[0]
			} else {
				product[i] = pool[this.indices[i]]
				break
			}
		}
	}

	this.current = product
	this.iterations++
	return product
}

func (this *ProductIterator) HasNext() bool {
	return this.iterations < this.total
}

func (this *ProductIterator) EmptyProduct() []interface{} {
	return make([]interface{}, len(this.pools))
}

func (this *ProductIterator) Total() *big.Int {
	sizes := make([]int, len(this.pools))
	for i := range sizes {
		sizes[i] = len(this.pools[i])
	}

	return TotalProduct(sizes...)
}

func (this *ProductIterator) Reset() {
	this.iterations = 0
	this.total = this.Total().Int64()

	n := len(this.pools)
	this.indices = make([]int, n)
	this.current = make([]interface{}, n)
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
