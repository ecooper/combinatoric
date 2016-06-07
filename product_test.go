package combinatoric

import (
	"fmt"
	"math/big"
	"testing"
)

func TestLenProduct(t *testing.T) {
	tests := []struct {
		n []int
		e *big.Int
	}{
		{[]int{2, 2, 2}, big.NewInt(8)},
		{[]int{2, 1, 1, 5}, big.NewInt(10)},
	}

	for _, test := range tests {
		if v := LenProduct(test.n...).Int64(); v != test.e.Int64() {
			t.Errorf("LenProduct(%v) != %v, got %v", test.n, test.e, v)
		}
	}
}

func TestProduct(t *testing.T) {
	tests := []struct {
		v [][]interface{}
		e [][]interface{}
	}{
		{
			[][]interface{}{
				{"A", "B"},
				{"C", "D"},
				{"E"},
			},
			[][]interface{}{
				{"A", "C", "E"},
				{"A", "D", "E"},
				{"B", "C", "E"},
				{"B", "D", "E"},
			},
		},
	}

	for _, test := range tests {
		products, _ := Product(test.v)
		for i := 0; products.HasNext(); i++ {
			c := products.Next()
			if fmt.Sprint(c) != fmt.Sprint(test.e[i]) {
				t.Errorf("Got unexpected product, %v at %v. Expected %v", c, i, test.e)
			}
		}
	}
}

func BenchmarkProduct(b *testing.B) {
	pools := [][]interface{}{
		{"A", "B", "C"},
		{"A", "B"},
		{"A"},
	}
	results := make([][]interface{}, LenProduct(3, 2, 1).Int64())

	for i := 0; i < b.N; i++ {
		product, _ := Product(pools)
		for c := 0; product.HasNext(); c++ {
			results[c] = product.Next()
		}
	}
}
