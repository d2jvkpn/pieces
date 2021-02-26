package explore

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestLogTimeStr(t *testing.T) {
	fmt.Println(LogTime("0102T15:04:05"))
	fmt.Println(LogTime("2006-01-02T15:04:05"))
}

func TestFloorClock(t *testing.T) {
	fmt.Println(FloorClock(10, "S"))
	fmt.Println(FloorClock(5, "M"))
	fmt.Println(FloorClock(1, "H"))
	fmt.Println(FloorClock(5, "H"))
}

func TestPedrail_x1(t *testing.T) {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	pc, _ := NewPedrail(5, "S", "TestPedrail", nil)

	pc.Start()
	for i := 0; i < 60; i++ {
		time.Sleep(time.Second*time.Duration(rd.Int63n(5)) + 1)
		pc.Add(1)
		// fmt.Println(pc.DataJSON())
		// fmt.Println(pc.Keys())
	}

	pc.Stop()
	fmt.Println(pc.DataJSON())
}

func TestPedrail_x2(t *testing.T) {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	pc, _ := NewPedrail(2, "M", "TestPedrail", nil)

	pc.Start()
	for i := 0; i < 300; i++ {
		time.Sleep(time.Second*time.Duration(rd.Int63n(3)) + 1)
		pc.Add(1)
		fmt.Println(pc.DataJSON())
		// fmt.Println(pc.Keys())
	}

	pc.Stop()
	fmt.Println(pc.DataJSON())
}
