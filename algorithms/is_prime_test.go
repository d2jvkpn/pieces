package algorithms

import (
	"fmt"
	"testing"
)

func TestIsPrime(t *testing.T) {
	for i := 1; i <= 20; i++ {
		isprime, err := IsPrime(i)
		fmt.Printf(">>> %d, %t, %v\n", i, isprime, err)
	}
}
