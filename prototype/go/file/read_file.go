package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var (
		n    int
		ret  int64
		err  error
		file *os.File
	)

	if file, err = os.Open("example.txt"); err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 2)
	n, err = file.ReadAt(buf, int64(2))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ReadAt: [%d] %d: %s\n", n, len(buf), buf)

	if ret, err = file.Seek(10, io.SeekStart); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Seek: %d\n", ret)

	if ret, err = file.Seek(10, io.SeekCurrent); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Seek: %d\n", ret)

	if ret, err = file.Seek(10, io.SeekEnd); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Seek: %d\n", ret)
}
