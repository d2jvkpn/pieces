package rover

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeS2N(t *testing.T) {
	var (
		s   string
		err error
		now time.Time
	)

	now = time.Now()

	if s, err = TimeConvert("t2n", now.Format(DefaultTimeFormat)); err != nil {
		t.Error(err)
	}
	println(s)

	if s, err = TimeConvert("t2u", now.Format(DefaultTimeFormat)); err != nil {
		t.Error(err)
	}
	println(s)

	if s, err = TimeConvert("t2m", now.Format(DefaultTimeFormat)); err != nil {
		t.Error(err)
	}
	println(s)

	if s, err = TimeConvert("m2t", "1574868339000"); err != nil {
		t.Error(err)
	}
	println(s)

	if s, err = TimeConvert("u2t", "1574868339000000"); err != nil {
		t.Error(err)
	}
	println(s)

}

////
func TestTimeSplitter_x1(t *testing.T) {
	start := time.Now()
	end := start.Add(time.Hour * 3)
	end = end.Add(time.Minute * 17)
	tsp, _ := NewTimeSplitter(start, end, time.Minute*39)

	fmt.Println(tsp.Start)
	fmt.Println(tsp.GetNK())

	var s, e time.Time
	var err error
	var n, k int64

	for {
		n, k = tsp.GetNK()
		s, e, err = tsp.Next()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("-->", n, k)
		fmt.Println("   ", s)
		fmt.Println("   ", e)
	}

	fmt.Println(tsp.End)
}
