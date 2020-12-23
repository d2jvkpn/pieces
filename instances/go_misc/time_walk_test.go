package rover

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeWalk_x1(t *testing.T) {
	now := time.Now()

	var unites []string = []string{"minute", "hour", "day", "week", "month",
		"season", "year"}

	fmt.Println("Now:", now)
	for _, u := range unites {
		nt, _ := TimeFloor(now, u)
		fmt.Println("TimeFloor", u, nt)
		nt, _ = TimeCeil(now, u)
		fmt.Println("TimeCeil", u, nt)
	}

}

func TestTimeWalk_x2(t *testing.T) {
	now := time.Now()
	var unites []string = []string{"M", "H", "d", "w", "m", "s", "y"}
	var t0, t1 time.Time

	fmt.Println("Now:", now)
	for _, u := range unites {
		t0, t1, _ = TimeScaleRange(now, fmt.Sprintf(" -2 %s ", u), TimeFloor)
		fmt.Println("TimeFloor -2", u, t0, t1)
		t0, t1, _ = TimeScaleRange(now, fmt.Sprintf("2%s", u), TimeFloor)
		fmt.Println("TimeFloor 2", u, t0, t1)
		t0, t1, _ = TimeScaleRange(now, fmt.Sprintf("-2%s", u), TimeCeil)
		fmt.Println("TimeCeil -2", u, t0, t1)
		t0, t1, _ = TimeScaleRange(now, fmt.Sprintf("2%s", u), TimeCeil)
		fmt.Println("TimeCeil 2", u, t0, t1)
	}
}

func TestParseDuration(t *testing.T) {
	fmt.Println(ParseNTU(" 5 d "))
	fmt.Println(ParseNTU("3m"))
}
