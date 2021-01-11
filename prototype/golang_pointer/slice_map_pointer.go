package main

import (
	"fmt"
)

func main() {
	dts := make([]Dt, 3)
	mp := make(map[int64]*Dt, 3)
	for i := range dts {
		dts[i].Id = int64(i + 1)
		// x := dts[i]
		// mp[dts[i].Id] = &x
		mp[dts[i].Id] = &dts[i]
	}

	n := 0
	do := func() {
		n++
		fmt.Println(">>>", n)
		fmt.Printf("%#v: \n    ", dts)
		for i := range dts {
			fmt.Printf("%p ", &dts[i])
		}
		fmt.Println()
		fmt.Printf("%v\n    %#v, %#v, %#v\n", mp, mp[1], mp[2], mp[3])
	}

	//// 1
	do()

	//// 2
	dts = dts[:0]
	do()

	//// 3
	for i := 4; i < 7; i++ {
		dts = append(dts, Dt{Id: int64(i)})
	}
	do()

	//// 4
	dts[0].Id = 9
	do()

	//// 5
	d := Dt{Id: 7}
	dts = append(dts, d)
	do()

	//// 6
	dts[0] = d
	do()
}

type Dt struct {
	Id int64
}
