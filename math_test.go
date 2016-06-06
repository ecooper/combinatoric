package combinatoric

import (
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
