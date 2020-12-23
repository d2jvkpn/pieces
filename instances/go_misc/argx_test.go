package rover

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestNewArgx(t *testing.T) {
	argx, _ := NewArgx()
	// fmt.Printf("%[1]T:\n    %#[1]v\n", argx.Start)
	// fmt.Printf("%[1]T:\n    %#[1]v\n", argx.Start)
	JsonTo(argx.Start, os.Stdout, true)
	argx.Log("TestA", "a", "b", 3, nil)
	argx.SetFile("test_data/xx.log")
	argx.Log("TestA", "c", 0b110100, nil, 4, "hello")
	fmt.Println(argx.file == nil)
	argx.Log("", "d", 0b110100, "qwert", 5, "world")
	time.Sleep(5 * time.Second)
	argx.LogBlock("TestA", "Hello, world!\nROVER")
	argx.Done(200, "OK")
	// fmt.Printf("%[1]T:\n    %#[1]v\n", argx.End)
	JsonTo(argx, os.Stdout, true)
	JsonToStdout(argx.End)
}
