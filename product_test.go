package combinatoric

import (
	"fmt"
	"testing"
)

func TestLenProduct(t *testing.T) {
	tests := []struct {
		n []int
		e uint64
	}{
		{[]int{2, 2, 2}, 8},
		{[]int{2, 1, 1, 5}, 10},
	}

	for _, test := range tests {
		if v := LenProduct(test.n...); v != test.e {
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
		{"A", "B", "C", "D", "E"},
		{"F", "G", "H"},
		{"I"},
	}
	results := make([][]interface{}, LenProduct(5, 3, 1))

	for i := 0; i < b.N; i++ {
		product, _ := Product(pools)
		for c := 0; product.HasNext(); c++ {
			results[c] = product.Next()
		}
	}
}
