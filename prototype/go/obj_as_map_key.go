package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
	mp := make(map[Data]bool, 1)

	d := Data{Ip: "127.0.0.1", Port: 1991}
	mp[d] = true

	fmt.Printf("%#v\n", d)
}

type Data struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}
