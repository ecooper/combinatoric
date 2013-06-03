package combinatoric

import (
	"math/big"
)

func factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, factorial(n.Sub(x, n)))
}

type CombinationIterator struct {
	pool    []interface{}
	r       int
	indices []int
	last    int
}

func (this *CombinationIterator) Next() []interface{} {
	combination := this.EmptyCombination()
	if this.last < 0 {
		this.last = 0
		for i := 0; i < this.r; i++ {
			combination[i] = this.pool[i]
		}
	} else {
		this.indices[this.last] = this.indices[this.last] + 1
		for j := this.last + 1; j < this.r; j++ {
			this.indices[j] = this.indices[j-1] + 1
		}

		for j := 0; j < this.r; j++ {
			combination[j] = this.pool[this.indices[j]]
		}
	}

	return combination
}

func (this *CombinationIterator) HasNext() bool {
	if this.r > len(this.pool) {
		return false
	}

	if this.last == -1 {
		return true
	}

	for i := this.r - 1; i > -1; i-- {
		this.last = i
		if this.indices[i] != this.last+len(this.pool)-this.r {
			return true
		}
	}

	return false
}

func (this *CombinationIterator) EmptyCombination() []interface{} {
	return make([]interface{}, this.r)
}

func (this *CombinationIterator) Reset() {
	this.last = -1
	this.indices = make([]int, this.r)
	for i := 0; i < this.r; i++ {
		this.indices[i] = i
	}

}

func TotalCombinations(n int, r int) (total *big.Int) {
	total = new(big.Int)
	d := new(big.Int)
	n64 := int64(n)
	r64 := int64(r)

	if n < r {
		total.SetInt64(0)
		return
	}

	d.Set(factorial(big.NewInt(n64 - r64)))
	d.Mul(d, factorial(big.NewInt(r64)))

	total.Set(factorial(big.NewInt(n64)))
	total.Div(total, d)

	return
}

func Combinations(pool []interface{}, r int) *CombinationIterator {
	combinations := new(CombinationIterator)

	combinations.pool = pool
	combinations.r = r
	combinations.Reset()

	return combinations
}
