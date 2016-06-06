package combinatoric

import (
	"math/big"
)

type PermutationIterator struct {
	pool    []interface{}
	r       int
	cycles  []int
	indices []int
	started bool
}

func (iter *PermutationIterator) Next() []interface{} {
	n := len(iter.pool)

	if iter.started == true {
		for i := iter.r - 1; i > -1; i-- {
			iter.cycles[i] -= 1
			if iter.cycles[i] == 0 {
				for x := i; x < n-1; x++ {
					v := iter.indices[x]
					iter.indices[x] = iter.indices[x+1]
					iter.indices[x+1] = v
				}
				iter.cycles[i] = n - i
			} else {
				x := iter.indices[i]
				iter.indices[i] = iter.indices[n-iter.cycles[i]]
				iter.indices[n-iter.cycles[i]] = x
				break
			}
		}
	} else {
		iter.started = true
	}

	permutation := iter.EmptyPermutation()
	for i := 0; i < iter.r; i++ {
		permutation[i] = iter.pool[iter.indices[i]]
	}

	return permutation
}

func (iter *PermutationIterator) HasNext() bool {
	if len(iter.pool) == iter.r && iter.started {
		return false
	}

	if iter.cycles[0] != 1 {
		return true
	}

	cf := 0
	for i := range iter.cycles {
		cf += iter.cycles[i]
	}

	return cf > iter.r
}

func (iter *PermutationIterator) EmptyPermutation() []interface{} {
	return make([]interface{}, iter.r)
}

func (iter *PermutationIterator) Reset() {
	n := len(iter.pool)
	iter.started = false

	iter.indices = make([]int, n)
	for i := range iter.indices {
		iter.indices[i] = i
	}

	iter.cycles = make([]int, n-(n-iter.r))
	for i := range iter.cycles {
		iter.cycles[i] = n - i
	}
}

func Permutations(pool []interface{}, r int) *PermutationIterator {
	permutations := new(PermutationIterator)

	permutations.pool = pool
	permutations.r = r
	permutations.Reset()

	return permutations
}

func TotalPermutations(n int, r int) *big.Int {
	n64 := int64(n)
	r64 := int64(r)

	if n < r {
		return big.NewInt(0)
	}

	total := big.NewInt(0)

	total.Div(
		factorial(big.NewInt(n64)),
		factorial(big.NewInt(n64-r64)),
	)

	return total
}
