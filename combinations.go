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
	pool      []interface{}
	r         int
	indices   []int
	completed int
	total     int
}

func (this *CombinationIterator) Next() []interface{} {
	if this.completed > 0 {
		i := this.r - 1
		for ; i > -1; i-- {
			if this.indices[i] != i+len(this.pool)-this.r {
				break
			}
		}
		if i > -1 {
			this.indices[i] += 1
			for j := i + 1; j < this.r; j++ {
				this.indices[j] = this.indices[j-1] + 1
			}

			this.completed += 1
		}
	} else {
		this.completed = 1
	}

	combination := this.EmptyCombination()
	for j := 0; j < this.r; j++ {
		combination[j] = this.pool[this.indices[j]]
	}

	return combination
}

func (this *CombinationIterator) HasNext() bool {
	if this.r > len(this.pool) {
		return false
	}

	return this.completed < this.total
}

func (this *CombinationIterator) EmptyCombination() []interface{} {
	return make([]interface{}, this.r)
}

func (this *CombinationIterator) Reset() {
	n := len(this.pool)
	if this.r > 1 {
		z := n - this.r + 1
		this.total = (n * z) / 2
	} else {
		this.total = n
	}

	this.completed = -1

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
