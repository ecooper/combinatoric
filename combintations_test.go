package combinatoric

import (
	"fmt"
	"math/big"
	"testing"
)

func TestLenCombinations(t *testing.T) {
	tests := []struct {
		n int
		r int
		e *big.Int
	}{
		{5, 3, big.NewInt(10)},
		{5, 5, big.NewInt(1)},
		{5, 6, big.NewInt(0)},
	}

	for _, test := range tests {
		if v := LenCombinations(test.n, test.r).Int64(); v != test.e.Int64() {
			t.Errorf("LenCombinations(%v, %v) != %v, got %v", test.n, test.r, test.e, v)
		}
	}
}

func TestCombinations(t *testing.T) {
	tests := []struct {
		v []interface{}
		r int
		e [][]interface{}
	}{
		{
			[]interface{}{"A", "B"},
			1,
			[][]interface{}{
				{"A"},
				{"B"},
			},
		},
		{
			[]interface{}{"A", "B", "C", "D"},
			2,
			[][]interface{}{
				{"A", "B"},
				{"A", "C"},
				{"A", "D"},
				{"B", "C"},
				{"B", "D"},
				{"C", "D"},
			},
		},
		{
			[]interface{}{"A", "B", "C", "D"},
			4,
			[][]interface{}{
				{"A", "B", "C", "D"},
			},
		},
		{
			[]interface{}{"A", "B", "C", "D", "E"},
			4,
			[][]interface{}{
				{"A", "B", "C", "D"},
				{"A", "B", "C", "E"},
				{"A", "B", "D", "E"},
				{"A", "C", "D", "E"},
				{"B", "C", "D", "E"},
			},
		},
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
			[]interface{}{"A", "B", "C", "D", "E"},
			1,
			[][]interface{}{
				{"A"},
				{"B"},
				{"C"},
				{"D"},
				{"E"},
			},
		},
		{
			[]interface{}{"A", "B", "C", "D", "E", "F", "G", "H", "I"},
			2,
			[][]interface{}{
				{"A", "B"},
				{"A", "C"},
				{"A", "D"},
				{"A", "E"},
				{"A", "F"},
				{"A", "G"},
				{"A", "H"},
				{"A", "I"},
				{"B", "C"},
				{"B", "D"},
				{"B", "E"},
				{"B", "F"},
				{"B", "G"},
				{"B", "H"},
				{"B", "I"},
				{"C", "D"},
				{"C", "E"},
				{"C", "F"},
				{"C", "G"},
				{"C", "H"},
				{"C", "I"},
				{"D", "E"},
				{"D", "F"},
				{"D", "G"},
				{"D", "H"},
				{"D", "I"},
				{"E", "F"},
				{"E", "G"},
				{"E", "H"},
				{"E", "I"},
				{"F", "G"},
				{"F", "H"},
				{"F", "I"},
				{"G", "H"},
				{"G", "I"},
				{"H", "I"},
			},
		},
	}

	for _, test := range tests {
		combinations, _ := Combinations(test.v, test.r)
		i := 0
		for ; combinations.HasNext(); i++ {
			c := combinations.Next()
			if fmt.Sprint(c) != fmt.Sprint(test.e[i]) {
				t.Errorf("Got unexpected combination, %v at %v. Expected %v", c, i, test.e[i])
			}
		}

		if int(len(test.e)) != int(i) {
			t.Errorf("Not enough combinations: %s, expected %s", i, len(test.e))
		}
	}
}

func BenchmarkCombination(b *testing.B) {
	pool := []interface{}{"A", "B", "C", "D", "E"}
	r := 2
	results := make([][]interface{}, LenCombinations(len(pool), r).Int64())

	for i := 0; i < b.N; i++ {
		combinations, _ := Combinations(pool, r)
		for c := 0; combinations.HasNext(); c++ {
			results[c] = combinations.Next()
		}
	}
}
