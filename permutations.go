package combinatoric

import (
	"log"
	"math/big"
)

type PermutationIterator struct {
	pool       []interface{}
	r          int
	cycles     []int
	indices    []int
	total      int64
	iterations int64
}

func (this *PermutationIterator) Next() []interface{} {
	n := len(this.pool)

	if this.iterations > 0 {
		log.Print(this.cycles, this.indices)
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
				break
			}
		}
	}

	permutation := this.EmptyPermutation()
	for i := 0; i < this.r; i++ {
		permutation[i] = this.pool[this.indices[i]]
	}

	this.iterations++
	return permutation
}

func (this *PermutationIterator) HasNext() bool {
	return this.iterations < this.total
}

func (this *PermutationIterator) EmptyPermutation() []interface{} {
	return make([]interface{}, this.r)
}

func (this *PermutationIterator) Reset() {
	n := len(this.pool)
	this.total = TotalPermutations(n, this.r).Int64()
	this.iterations = 0

	this.indices = make([]int, n)
	for i := range this.indices {
		this.indices[i] = i
	}

	this.cycles = make([]int, n-this.r)
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
