package main

import (
	"fmt"
	"os"
	"testing"
)

func TestJsonTo(t *testing.T) {
	data, _ := NewArgx()
	JsonTo(data, os.Stdout, true)
	JsonToFile(data, "test_data/argx.json", false)
}

func TestPrintData(t *testing.T) {
	d, _ := NewArgx()
	PrintData(d)
}

func TestSaveData(t *testing.T) {
	d, _ := NewArgx()
	SaveData(d, "test_data/argx.json")
}

func TestJsonToStdout(t *testing.T) {
	d1 := make([]string, 10)

	JsonToStdout(d1)

	d2 := 10
	err := JsonToStdout(d2)
	fmt.Println(err)
}

func TestJsonStr(t *testing.T) {
	d1 := make(map[int64]int64, 1)
	d1[1212812] = 5
	fmt.Println(JsonStr(d1, false))
}
