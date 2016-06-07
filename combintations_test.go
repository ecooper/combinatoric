package combinatoric

import (
	"reflect"
	"testing"
)

func TestLenCombinations(t *testing.T) {
	tests := []struct {
		n int
		r int
		e uint64
	}{
		{5, 3, 10},
		{5, 5, 1},
		{5, 6, 0},
	}

	for _, test := range tests {
		if v := lenCombinations(test.n, test.r); v != test.e {
			t.Errorf("lenCombinations(%v, %v) != %v, got %v", test.n, test.r, test.e, v)
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
			if !reflect.DeepEqual(c, test.e[i]) {
				t.Errorf("Got unexpected combination, %v at %v. Expected %v", c, i, test.e[i])
			}
		}

		if len(test.e) != i {
			t.Errorf("Not enough combinations: %d, expected %d", i, len(test.e))
		}
	}
}

func TestCombinationsReset(t *testing.T) {
	combinations, _ := Combinations([]interface{}{"A", "B", "C"}, 2)

	combinations.First()
	combinations.Reset()

	if combinations.iters != 0 {
		t.Error("iters should be zero after reset")
	}

	if !reflect.DeepEqual(combinations.indices, []int{0, 1}) {
		t.Errorf("indicies not at starting values, expected %d got %d", []int{0, 1}, combinations.indices)
	}
}

func BenchmarkCombination(b *testing.B) {
	pool := []interface{}{"A", "B", "C", "D", "E"}
	r := 5
	results := make([][]interface{}, lenCombinations(len(pool), r))

	for i := 0; i < b.N; i++ {
		combinations, _ := Combinations(pool, r)
		for c := 0; combinations.HasNext(); c++ {
			results[c] = combinations.Next()
		}
		combinations.Reset()
	}
}
