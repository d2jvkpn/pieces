package explore

import (
	"fmt"
)

func MergeSort(slice []int) (out []int) {
	if len(slice) < 2 {
		return slice
	}

	m := (len(slice) + 1) / 2
	s1, s2 := MergeSort(slice[:m]), MergeSort(slice[m:]) // recurrsion
	fmt.Printf("    ~~~ s1 = %v, s2 = %v\n", s1, s2)
	out = make([]int, 0, len(slice))

	for i, j := 0, 0; ; {
		if i == len(s1) || j == len(s2) {
			out = append(out, s2[j:]...) // when i == len(s1)
			out = append(out, s1[i:]...) // when j == len(s2)
			break
		}

		if s1[i] < s2[j] {
			out = append(out, s1[i])
			i++
		} else {
			out = append(out, s2[j])
			j++
		}
	}

	return
}

func InstMergeSort1() {
	fmt.Println(">>> InstMergeSort:")
	slice := []int{14, 33, 10, 27, 19, 35, 42, 44}
	fmt.Printf("    slice = %v\n", slice)

	out := MergeSort(slice)
	fmt.Printf("    out = %v\n", out)
}

func InstMergeSort2() {
	fmt.Println(">>> InstMergeSort2:")
	slice := []int{14, 33, 10, 27, 19, 35, 42, 44, 18}
	fmt.Printf("    slice = %v\n", slice)

	out := MergeSort(slice)
	fmt.Printf("    out = %v\n", out)
}
