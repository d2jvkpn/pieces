package main

import (
	"embed"
	"fmt"
	"log"
)

//go:embed fs
var Dir embed.FS

func main() {
	dirs, err := Dir.ReadDir("fs")
	if err != nil {
		log.Fatal(err)
	}

	for i := range dirs {
		fmt.Println(">>> found:", dirs[i].Name())
	}
}
