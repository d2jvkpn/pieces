package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"time"
)

// marshal data to json([]byte), and write to an io.Writer,
func JsonTo(data interface{}, out io.Writer, readable bool) (err error) {
	if out == nil {
		err = fmt.Errorf("out io.Writer is nil")
		return
	}

	// intead json.Marshal to []byte and write []byte to out(io.Writer)
	encoder := json.NewEncoder(out)
	if readable {
		encoder.SetIndent("", "  ")
	}

	if err = encoder.Encode(data); err != nil {
		return
	}

	return
}

// masharshal data to json([]byte) and save to file, parents directories of file
// will be created
func JsonToFile(data interface{}, out string, readable bool) (err error) {
	var file *os.File

	if out == "-" || out == "." || out == "" {
		JsonTo(data, os.Stdout, true)
		return
	}

	if err = os.MkdirAll(path.Dir(out), 0755); err != nil {
		return
	}

	if file, err = os.Create(out); err != nil {
		return
	}
	defer file.Close()

	if err = JsonTo(data, file, readable); err != nil {
		return
	}

	return
}

// masharshal data to json([]byte) to os.Stdout
func JsonToStdout(data interface{}) (err error) {
	if data == nil {
		err = fmt.Errorf("input data is nil")
		return
	}

	err = JsonTo(data, os.Stdout, true)

	return
}

// print data to stdout in json format
func PrintData(data interface{}) {
	var err error

	fmt.Fprintf(os.Stderr, "\n[%s] >>> Print data %T:\n",
		time.Now().Format(DefaultTimeFormat), data)

	if err = JsonToStdout(data); err != nil {
		fmt.Printf("%#v\n", data)
	}
}

// save data in json
func SaveData(d interface{}, p string) {
	var err error

	fmt.Fprintf(os.Stderr, "\n[%s] >>> Saved data %T:\n",
		time.Now().Format(DefaultTimeFormat), d)

	if err = JsonToFile(d, p, true); err == nil {
		fmt.Println("    saved", p)
	} else {
		fmt.Printf("    failed to save to %s: %v\n", p, err)
	}
}

func JsonStr(data interface{}, readable bool) (str string, err error) {
	var bts []byte
	if bts, err = json.Marshal(data); err != nil {
		return
	}
	str = string(bts)

	return
}
