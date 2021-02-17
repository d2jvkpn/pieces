package main

import (
	"embed"
	"fmt"
	"log"
)

//go:embed fs/hello.txt
var fs embed.FS

func main() {
	bts, err := fs.ReadFile("fs/hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(bts))
}
