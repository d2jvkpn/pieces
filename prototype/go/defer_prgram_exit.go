package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello, world!")

	defer func() {
		fmt.Println(">>> defer")
	}()

	return
	os.Exit(0)                      // not execute defer functons
	log.Fatal("unexpected error\n") // not execute defer functons
}
