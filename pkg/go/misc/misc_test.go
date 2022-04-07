package misc

import (
	"testing"
)

func TestListIndex(t *testing.T) {
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

			out := ListIndex(tt[i].List, tt[i].Item)
			if out != tt[i].Expected {
				t.Fatalf("exp: %d, got: %d\n", tt[i].Expected, out)
			}
		}

		t.Run(tt[i].Name, tf)
	}
}

func TestUniqList(t *testing.T) {
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

			out := UniqList(tt[i].List)
			if !EqualList(out, tt[i].Expected) {
				t.Fatalf("unepected: %v, %d\n", tt[i].Expected, out)
			}
		}

		t.Run(tt[i].Name, tf)
	}
}
