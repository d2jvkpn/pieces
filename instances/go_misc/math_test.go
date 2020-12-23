package rover

import (
	"fmt"
	"testing"
)

////
func TestGrowthRatioInPerc(t *testing.T) {
	fmt.Println(GrowthRatioInPerc(3, 5))

	fmt.Println(GrowthRatioInPerc(0, 9))
}

func TestDivmod(t *testing.T) {
	fmt.Println(Divmod(100, 3))
	fmt.Println(Divmod(100, 4))
	fmt.Println(Divmod(9, 12))
}
