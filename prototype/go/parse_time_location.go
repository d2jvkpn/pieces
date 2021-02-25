package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, world!")

	var templ string = "2006-01-02 15:04:05"

	ts := "2020-01-16 10:56:09"
	t1, _ := time.ParseInLocation(templ, ts, time.Local)
	t2, _ := time.Parse(templ, ts)

	fmt.Println(t1) // +08:00
	fmt.Println(t2) // +00:00
}
