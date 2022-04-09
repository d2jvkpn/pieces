package misc

import (
	"encoding/base64"
	"fmt"
	"path/filepath"
	"testing"
)

func TestVectorIndex(t *testing.T) {
	tt := []struct {
		Name     string
		List     []string
		Item     string
		Expected int
	}{
		{"x1", []string{"a", "b", "c"}, "a", 0},
		{"x2", []string{"a", "b", "c"}, "z", -1},
	}

	for i := range tt {
		i := i
		tf := func(t *testing.T) {
			t.Parallel()

			out := VectorIndex(tt[i].List, tt[i].Item)
			if out != tt[i].Expected {
				t.Fatalf("exp: %d, got: %d\n", tt[i].Expected, out)
			}
		}

		t.Run(tt[i].Name, tf)
	}
}

func TestUniqVector(t *testing.T) {
	tt := []struct {
		Name     string
		List     []byte
		Expected []byte
	}{
		{"x1", []byte{'a', 'b', 'a'}, []byte{'a', 'b'}},
		{"x2", []byte{'a', 'b', 'c'}, []byte{'a', 'b', 'c'}},
	}

	for i := range tt {
		i := i
		tf := func(t *testing.T) {
			t.Parallel()

			out := UniqVector(tt[i].List)
			if !EqualVector(out, tt[i].Expected) {
				t.Fatalf("unepected: %v, %d\n", tt[i].Expected, out)
			}
		}

		t.Run(tt[i].Name, tf)
	}
}

func TestBase64(t *testing.T) {
	// bts := []byte("Hello, world, 你好!")
	bts := []byte{0xff}
	out := Base64Encode(bts)
	fmt.Println(out, base64.StdEncoding.EncodeToString(bts))

	if bts, err := Base64Decode(out); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(string(bts))
	}
}

func TestFilepath(t *testing.T) {
	p := "a/b/c.tar.gz"
	fmt.Println(filepath.Base(p), filepath.Ext(p))
}

func TestFileSaveName(t *testing.T) {
	fmt.Println(FileSaveName("misc.go"))
}
