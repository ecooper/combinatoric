package combinatoric

import (
	"math/big"
)

var bigIntIncr *big.Int = big.NewInt(1)

func factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, factorial(n.Sub(x, n)))
}
