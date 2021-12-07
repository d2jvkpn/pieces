package interview

import (
	"fmt"
	"math"
)

func IsPrime(num int) (yes bool, err error) {
	if num < 2 {
		return false, fmt.Errorf("invlaid number")
	} else if num == 2 {
		return true, nil
	}

	r := math.Sqrt(float64(num))
	for i := 2; i <= int(r); i++ {
		if num%i == 0 {
			return false, nil
		}
	}

	return true, nil
}

func InstIsPrime() {
	fmt.Println(">>> InstIsPrime:")

	for i := 1; i <= 20; i++ {
		isprime, err := IsPrime(i)
		fmt.Printf(">>> %d, %t, %v\n", i, isprime, err)
	}
}
