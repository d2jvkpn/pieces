package rover

import (
	"fmt"
	"testing"
)

func TestBts2File(t *testing.T) {
	bts := []byte{'a', 'b', 'c', 'x', '\n'}

	err := Bts2File(bts, "test_data/abcx.txt")
	if err != nil {
		t.Fatal(err)
	}

	return
}

func TestColumns(t *testing.T) {
	a := [][]string{
		{"aaa", "BBBBBBBBB", "D"},
		{"xxx", "yyyyyyyyyyyyyyyyyyyyyyyyyy", "z"},
		{"1", "2", "3"},
	}

	Columns(a, 4, 2)
}

// ln -sr path_test.go ../a.go
// ln -sr ../a.go ../b.go
func TestPath(t *testing.T) {
	fmt.Println(DecomposePath("../a.go"))
	fmt.Println(RealPath("../b.go"))
}

func TestMD5(t *testing.T) {
	fmt.Println("Hello, world!")

	a := "aaa"
	A := "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
	b := "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"

	fmt.Println("a IsMD5:", IsMD5(a))
	fmt.Println("a IsMD5Upper:", IsMD5Upper(a))
	fmt.Println("a IsMD5Lower:", IsMD5Lower(a))

	fmt.Println("A IsMD5:", IsMD5(A))
	fmt.Println("A IsMD5Upper:", IsMD5Upper(A))
	fmt.Println("A IsMD5Lower:", IsMD5Lower(A))

	fmt.Println("b IsMD5:", IsMD5(b))
	fmt.Println("b IsMD5Upper:", IsMD5Upper(b))
	fmt.Println("b IsMD5Lower:", IsMD5Lower(b))
}
