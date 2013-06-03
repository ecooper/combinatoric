package combinatoric

import (
	"math/big"
)

type PermutationIterator struct {
	pool      []interface{}
	r         int
	cycles    []int
	indices   []int
	completed int
}

func (this *PermutationIterator) Next() []interface{} {
	n := len(this.pool)

	if this.completed > -1 {
		for i := this.r - 1; i > -1; i-- {
			this.cycles[i] = this.cycles[i] - 1
			if this.cycles[i] == 0 {
				for x := i; x < len(this.indices)-1; x++ {
					v := this.indices[x]
					this.indices[x] = this.indices[x+1]
					this.indices[x+1] = v
				}
				this.cycles[i] = n - i
			} else {
				x := this.indices[i]
				this.indices[i] = this.indices[n-this.cycles[i]]
				this.indices[n-this.cycles[i]] = x
				this.completed += i
				break
			}
		}
	} else {
		this.completed = 0
	}

	permutation := this.EmptyPermutation()
	for i := 0; i < this.r; i++ {
		permutation[i] = this.pool[this.indices[i]]
	}

	return permutation
}

func (this *PermutationIterator) HasNext() bool {
	return this.completed < len(this.pool)-this.r*this.r
}

func (this *PermutationIterator) EmptyPermutation() []interface{} {
	return make([]interface{}, this.r)
}

func (this *PermutationIterator) Reset() {
	n := len(this.pool)

	this.indices = make([]int, n)
	for i := range this.indices {
		this.indices[i] = i
	}

	this.cycles = make([]int, n+1-this.r)
	for i := range this.cycles {
		this.cycles[i] = n - i
	}
}

func Permutations(pool []interface{}, r int) *PermutationIterator {
	permutations := new(PermutationIterator)

	permutations.pool = pool
	permutations.r = r
	permutations.Reset()

	return permutations
}

func TotalPermutations(n int, r int) (total *big.Int) {
	total = new(big.Int)
	d := new(big.Int)
	n64 := int64(n)
	r64 := int64(r)

	if n < r {
		total.SetInt64(0)
		return
	}

	d.Set(factorial(big.NewInt(n64 - r64)))
	total.Set(factorial(big.NewInt(n64)))
	total.Div(total, d)

	return
}
