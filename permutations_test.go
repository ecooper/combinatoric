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
	}

	for _, test := range tests {
		permutations := Permutations(test.v, test.r)
		for i := 0; permutations.HasNext(); i++ {
			p := permutations.Next()
			if fmt.Sprint(p) != fmt.Sprint(test.e[i]) {
				t.Errorf("Got unexpected permutations, %v at %v. Expected %v", p, i, test.e[i])
			}
		}
	}
}
