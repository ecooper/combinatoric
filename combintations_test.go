package combinatoric

import (
	"fmt"
	"math/big"
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		v *big.Int
		e *big.Int
	}{
		{big.NewInt(5), big.NewInt(120)},
		{big.NewInt(1), big.NewInt(1)},
		{big.NewInt(0), big.NewInt(1)},
	}

	for _, test := range tests {
		if r := factorial(test.v).Int64(); r != test.e.Int64() {
			t.Errorf("factorial(%v) != %v, got %v", test.v, test.e, r)
		}
	}
}

func TestTotalCombinations(t *testing.T) {
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
		if v := TotalCombinations(test.n, test.r).Int64(); v != test.e.Int64() {
			t.Errorf("TotalCombinations(%v, %v) != %v, got %v", test.n, test.r, test.e, v)
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
	}

	for _, test := range tests {
		combinations := Combinations(test.v, test.r)
		for i := 0; combinations.HasNext(); i++ {
			c := combinations.Next()
			if fmt.Sprint(c) != fmt.Sprint(test.e[i]) {
				t.Errorf("Got unexpected combination, %v at %v. Expected %v", c, i, test.e)
			}
		}
	}
}
