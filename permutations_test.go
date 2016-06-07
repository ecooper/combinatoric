package combinatoric

import (
	"fmt"
	"math/big"
	"testing"
)

func TestTotalPermutations(t *testing.T) {
	tests := []struct {
		n int
		r int
		e *big.Int
	}{
		{5, 2, big.NewInt(20)},
		{5, 1, big.NewInt(5)},
		{5, 6, big.NewInt(0)},
	}

	for _, test := range tests {
		if v := TotalPermutations(test.n, test.r).Int64(); v != test.e.Int64() {
			t.Errorf("TotalCombinations(%v, %v) != %v, got %v", test.n, test.r, test.e, v)
		}
	}
}

func TestPermutations(t *testing.T) {
	tests := []struct {
		v []interface{}
		r int
		e [][]interface{}
	}{
		{
			[]interface{}{"A", "B", "C"},
			1,
			[][]interface{}{
				{"A"},
				{"B"},
				{"C"},
			},
		},
		{
			[]interface{}{"A", "B", "C", "D"},
			4,
			[][]interface{}{
				{"A", "B", "C", "D"},
				{"A", "B", "D", "C"},
				{"A", "C", "B", "D"},
				{"A", "C", "D", "B"},
				{"A", "D", "B", "C"},
				{"A", "D", "C", "B"},
				{"B", "A", "C", "D"},
				{"B", "A", "D", "C"},
				{"B", "C", "A", "D"},
				{"B", "C", "D", "A"},
				{"B", "D", "A", "C"},
				{"B", "D", "C", "A"},
				{"C", "A", "B", "D"},
				{"C", "A", "D", "B"},
				{"C", "B", "A", "D"},
				{"C", "B", "D", "A"},
				{"C", "D", "A", "B"},
				{"C", "D", "B", "A"},
				{"D", "A", "B", "C"},
				{"D", "A", "C", "B"},
				{"D", "B", "A", "C"},
				{"D", "B", "C", "A"},
				{"D", "C", "A", "B"},
				{"D", "C", "B", "A"},
			},
		},
		{
			[]interface{}{"A", "B", "C", "D"},
			2,
			[][]interface{}{
				{"A", "B"},
				{"A", "C"},
				{"A", "D"},
				{"B", "A"},
				{"B", "C"},
				{"B", "D"},
				{"C", "A"},
				{"C", "B"},
				{"C", "D"},
				{"D", "A"},
				{"D", "B"},
				{"D", "C"},
			},
		},
		{
			[]interface{}{"A", "B", "C", "D", "E"},
			3,
			[][]interface{}{
				{"A", "B", "C"},
				{"A", "B", "D"},
				{"A", "B", "E"},
				{"A", "C", "B"},
				{"A", "C", "D"},
				{"A", "C", "E"},
				{"A", "D", "B"},
				{"A", "D", "C"},
				{"A", "D", "E"},
				{"A", "E", "B"},
				{"A", "E", "C"},
				{"A", "E", "D"},
				{"B", "A", "C"},
				{"B", "A", "D"},
				{"B", "A", "E"},
				{"B", "C", "A"},
				{"B", "C", "D"},
				{"B", "C", "E"},
				{"B", "D", "A"},
				{"B", "D", "C"},
				{"B", "D", "E"},
				{"B", "E", "A"},
				{"B", "E", "C"},
				{"B", "E", "D"},
				{"C", "A", "B"},
				{"C", "A", "D"},
				{"C", "A", "E"},
				{"C", "B", "A"},
				{"C", "B", "D"},
				{"C", "B", "E"},
				{"C", "D", "A"},
				{"C", "D", "B"},
				{"C", "D", "E"},
				{"C", "E", "A"},
				{"C", "E", "B"},
				{"C", "E", "D"},
				{"D", "A", "B"},
				{"D", "A", "C"},
				{"D", "A", "E"},
				{"D", "B", "A"},
				{"D", "B", "C"},
				{"D", "B", "E"},
				{"D", "C", "A"},
				{"D", "C", "B"},
				{"D", "C", "E"},
				{"D", "E", "A"},
				{"D", "E", "B"},
				{"D", "E", "C"},
				{"E", "A", "B"},
				{"E", "A", "C"},
				{"E", "A", "D"},
				{"E", "B", "A"},
				{"E", "B", "C"},
				{"E", "B", "D"},
				{"E", "C", "A"},
				{"E", "C", "B"},
				{"E", "C", "D"},
				{"E", "D", "A"},
				{"E", "D", "B"},
				{"E", "D", "C"},
			},
		},
	}

	for _, test := range tests {
		permutations := Permutations(test.v, test.r)
		i := 0
		for ; permutations.HasNext(); i++ {
			p := permutations.Next()
			if fmt.Sprint(p) != fmt.Sprint(test.e[i]) {
				t.Errorf("Got unexpected permutations, %v at %v. Expected %v", p, i, test.e[i])
			}
		}

		if int(len(test.e)) != int(i) {
			t.Errorf("Not enough combinations: %s, expected %s", i, len(test.e))
		}
	}
}

func BenchmarkPermutation(b *testing.B) {
	pool := []interface{}{"A", "B", "C", "D", "E"}
	r := 2
	results := make([][]interface{}, TotalPermutations(len(pool), r).Int64())

	for i := 0; i < b.N; i++ {
		permutations := Permutations(pool, r)
		for c := 0; permutations.HasNext(); c++ {
			results[c] = permutations.Next()
		}
	}
}
