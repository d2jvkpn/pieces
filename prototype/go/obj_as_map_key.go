package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
	mp := make(map[Data]bool, 1)

	d := Data{Ip: "127.0.0.1", Port: 1991}
	d.Pointer = new(int)
	mp[d] = true

	fmt.Printf("%#v\n", d)

	inMap := func() (ok bool) {
		_, ok = mp[d]
		return ok
	}

	fmt.Println(inMap()) // true

	*(d.Pointer) = 10
	fmt.Println(inMap()) // true

	d.Pointer = nil
	fmt.Println(inMap()) // false
}

type Data struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
	// E    func() string // invalid map key type Data
	Pointer *int `json:"pointer"`
}

func (d *Data) Echo() string {
	return "hello"
}
