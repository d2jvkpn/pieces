package main

import (
	"fmt"
)

func main() {
	fmt.Println(SieveOfEratosthenes(100))
}

func SieveOfEratosthenes(n int) []int {
	// Finds all primes up to n
	list := make([]bool, n+1)
	primes := make([]int, 0)

	for i := 2; i < n+1; i++ {
		list[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if list[p] {
			for i := p * 2; i <= n; i += p {
				list[i] = false
			}
		}
	}

	for p := 2; p <= n; p++ {
		if list[p] {
			primes = append(primes, p)
		}
	}

	return primes
}
