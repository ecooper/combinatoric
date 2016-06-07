package combinatoric

import (
	"math/big"
)

type PermutationIterator struct {
	pool    []interface{}
	n       int
	r       int
	cycles  []int
	indices []int

	max   *big.Int
	iters *big.Int

	res []interface{}
}

func (iter *PermutationIterator) Next() []interface{} {
	if iter.res[0] != nil {
		for i := iter.r - 1; i > -1; i-- {
			iter.cycles[i] -= 1
			if iter.cycles[i] == 0 {
				for x := i; x < iter.n-1; x++ {
					v := iter.indices[x]
					iter.indices[x] = iter.indices[x+1]
					iter.indices[x+1] = v
				}
				iter.cycles[i] = iter.n - i
			} else {
				x := iter.indices[i]
				iter.indices[i] = iter.indices[iter.n-iter.cycles[i]]
				iter.indices[iter.n-iter.cycles[i]] = x
				break
			}
		}
	}

	for i := 0; i < iter.r; i++ {
		iter.res[i] = iter.pool[iter.indices[i]]
	}

	iter.iters.Add(iter.iters, bigIntIncr)

	return iter.res
}

func (iter *PermutationIterator) HasNext() bool {
	return iter.iters.Cmp(iter.max) == -1
}

func (iter *PermutationIterator) Reset() {
	iter.iters = big.NewInt(0)
	iter.max = iter.Len()

	iter.indices = make([]int, iter.n)
	for i := range iter.indices {
		iter.indices[i] = i
	}

	iter.cycles = make([]int, iter.n-(iter.n-iter.r))
	for i := range iter.cycles {
		iter.cycles[i] = iter.n - i
	}
}

func (iter *PermutationIterator) Len() *big.Int {
	return iter.max
}

func LenPermutations(n int, r int) *big.Int {
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

func Permutations(pool []interface{}, r int) *PermutationIterator {
	iter := &PermutationIterator{
		pool:  pool,
		n:     len(pool),
		r:     r,
		res:   make([]interface{}, r, r),
		iters: big.NewInt(0),
		max:   LenPermutations(len(pool), r),
	}

	iter.Reset()

	return iter
}
