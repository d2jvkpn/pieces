package algorithms

import (
	"fmt"
	"testing"
)

func TestDeliveryBalance(t *testing.T) {
	m := []int64{2, 5, 3, 12, 1, 7, 8}

	fmt.Println(m)

	DeliveryBalance(m, 20)

	fmt.Println(m)
}
