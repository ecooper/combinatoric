package combinatoric

import (
	"reflect"
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
		if v := lenProduct(test.n...); v != test.e {
			t.Errorf("lenProduct(%v) != %v, got %v", test.n, test.e, v)
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
		i := 0

		for p := products.First(); p != nil; p = products.Next() {
			if !reflect.DeepEqual(p, test.e[i]) {
				t.Errorf("Got unexpected product, %v at %v. Expected %v", p, i, test.e)
			}
			i += 1
		}

		if len(test.e) != i {
			t.Errorf("Not enough products: %d, expected %d", i, len(test.e))
		}
	}
}

func TestProductsReset(t *testing.T) {
	products, _ := Product([][]interface{}{{"A", "B"}, {"C", "D"}})

	products.First()
	products.Reset()

	if products.iters != 0 {
		t.Error("iters should be zero after reset")
	}

	if !reflect.DeepEqual(products.indices, []int{0, 0}) {
		t.Errorf("indicies not at starting values, expected %d got %d", []int{0, 0}, products.indices)
	}
}

func BenchmarkProduct(b *testing.B) {
	pools := [][]interface{}{
		{"A", "B", "C", "D", "E"},
		{"F", "G", "H"},
		{"I"},
	}
	results := make([][]interface{}, lenProduct(5, 3, 1))

	for i := 0; i < b.N; i++ {
		product, _ := Product(pools)
		for c := 0; product.HasNext(); c++ {
			results[c] = product.Next()
		}
	}
}
