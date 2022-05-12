// $ go run 42.go
package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Printf("1/42 = %.8f, 1/24 = %.8f\n", 1.0/42.0, 1.0/24.0)

	var (
		e  = big.NewInt(3)
		v1 = big.NewInt(-80538738812075974)
		v2 = big.NewInt(80435758145817515)
		v3 = big.NewInt(12602123297335631)
	)

	v1.Exp(v1, e, nil)
	v2.Exp(v2, e, nil)
	v3.Exp(v3, e, nil)
	v1.Add(v1, v2).Add(v1, v3)

	fmt.Printf(
		"Life, the Universe and Everything: %d, %d, %d, %d\n",
		int(0b101010), int(0x2a), int('*'), v1,
	)
}
