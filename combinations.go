package combinatoric

import (
	"math/big"
)

type CombinationIterator struct {
	pool      []interface{}
	r         int
	indices   []int
	completed int
	total     int
}

func (iter *CombinationIterator) Next() []interface{} {
	if iter.completed > 0 {
		i := iter.r - 1
		for ; i > -1; i-- {
			if iter.indices[i] != i+len(iter.pool)-iter.r {
				break
			}
		}
		if i > -1 {
			iter.indices[i] += 1
			for j := i + 1; j < iter.r; j++ {
				iter.indices[j] = iter.indices[j-1] + 1
			}

			iter.completed += 1
		}
	} else {
		iter.completed = 1
	}

	combination := iter.EmptyCombination()
	for j := 0; j < iter.r; j++ {
		combination[j] = iter.pool[iter.indices[j]]
	}

	return combination
}

func (iter *CombinationIterator) HasNext() bool {
	if iter.r > len(iter.pool) {
		return false
	}

	return iter.completed < iter.total
}

func (iter *CombinationIterator) EmptyCombination() []interface{} {
	return make([]interface{}, iter.r)
}

func (iter *CombinationIterator) Reset() {
	n := len(iter.pool)
	if iter.r == n {
		iter.total = 1
	} else if iter.r > 1 {
		z := n - iter.r + 1
		iter.total = (n * z) / 2
	} else {
		iter.total = n
	}

	iter.completed = -1

	iter.indices = make([]int, iter.r)
	for i := 0; i < iter.r; i++ {
		iter.indices[i] = i
	}

}

func TotalCombinations(n int, r int) *big.Int {
	n64 := int64(n)
	r64 := int64(r)

	if n < r {
		return big.NewInt(0)
	}

	total := big.NewInt(0)

	total.Mul(
		factorial(big.NewInt(n64-r64)),
		factorial(big.NewInt(r64)),
	)

	total.Div(factorial(big.NewInt(n64)), total)

	return total
}

func Combinations(pool []interface{}, r int) *CombinationIterator {
	combinations := new(CombinationIterator)

	combinations.pool = pool
	combinations.r = r
	combinations.Reset()

	return combinations
}
