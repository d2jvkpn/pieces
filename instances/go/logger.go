package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type Logger struct {
	file   *os.File
	mutex  sync.Mutex
	writer io.Writer
	sep    string
	idx    int64
	strb   *strings.Builder
}

func NewLogger() (logger *Logger) {
	logger = new(Logger)
	logger.strb = &strings.Builder{}
	logger.SetWriter(os.Stdout, ", ") // set default logger
	return
}

// set logger, writer and seperator
func (logger *Logger) SetWriter(w io.Writer, sep string) (err error) {
	if logger == nil {
		err = fmt.Errorf("logger is nil")
		return
	}

	if w == nil {
		err = fmt.Errorf("writer is nil")
		return
	}

	var str string
	str = time.Now().Format(DeafultTimeFormat)
	logger.mutex.Lock()
	logger.writer, logger.sep = w, sep
	logger.writer.Write([]byte(str + " ~~~ SET LOGGER\n"))

	logger.mutex.Unlock()
	return
}

func (logger *Logger) SetSep(sep string) (err error) {
	if logger == nil {
		err = fmt.Errorf("logger is nil")
		return
	}

	logger.mutex.Lock()
	logger.sep = sep
	logger.mutex.Unlock()
	return
}

func (logger *Logger) GetIndex() (idx int64) {
	if logger == nil {
		return
	}

	logger.mutex.Lock()
	idx = logger.idx
	logger.mutex.Unlock()
	return logger.idx
}

// write a message to logger
func (logger *Logger) Log(tag string, rs ...interface{}) (n int, err error) {
	if logger == nil {
		err = fmt.Errorf("logger is nil")
		return
	}

	if logger.writer == nil {
		err = fmt.Errorf("Argx.writer is nil")
		return
	}

	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	// time.RFC3339
	logger.idx++
	logger.strb.WriteString(time.Now().Format(DeafultTimeFormat))
	logger.strb.WriteString(fmt.Sprintf(" %d", logger.idx))
	if tag != "" {
		logger.strb.WriteString(" " + tag + ":")
	}

	if len(rs) > 0 {
		logger.strb.WriteString(" %v")
		for i := 0; i < len(rs)-1; i++ {
			logger.strb.WriteString(logger.sep + "%v")
		}
	}

	rec := strings.TrimSpace(fmt.Sprintf(logger.strb.String(), rs...))
	// rec = strings.Replace(rec, "\n", "\n    ", -1)
	n, err = logger.writer.Write([]byte(rec + "\n"))
	logger.strb.Reset()
	return
}

// write a message to logger
func (logger *Logger) LogBlock(tag string, msg string) (n int, err error) {
	if logger == nil {
		err = fmt.Errorf("logger is nil")
		return
	}

	if logger.writer == nil {
		err = fmt.Errorf("Argx.writer is nil")
		return
	}

	logger.mutex.Lock()
	defer logger.mutex.Unlock()

	// time.RFC3339
	logger.idx++
	logger.strb.WriteString(time.Now().Format(DeafultTimeFormat))
	logger.strb.WriteString(fmt.Sprintf(" %d", logger.idx))
	if tag != "" {
		logger.strb.WriteString(" " + tag + ":")
	}

	logger.strb.WriteString("\n    " +
		strings.Replace(strings.TrimSpace(msg), "\n", "\n    ", -1))

	// rec = strings.Replace(rec, "\n", "\n    ", -1)
	n, err = logger.writer.Write([]byte(logger.strb.String() + "\n"))
	logger.strb.Reset()
	return
}

func (logger *Logger) SetFile(fp string) (err error) {
	if logger == nil {
		err = fmt.Errorf("logger is nil")
		return
	}

	var file *os.File
	var str string

	if err = os.MkdirAll(filepath.Dir(fp), 0755); err != nil {
		return
	}

	if file, err = os.Create(fp); err != nil {
		return
	}

	str = time.Now().Format(DeafultTimeFormat)

	logger.mutex.Lock()
	if logger.file != nil {
		logger.file.Close()
	}

	logger.file = file
	logger.writer = logger.file
	logger.writer.Write([]byte(str + " ~~~ SET LOGGER file: " + fp + "\n"))

	logger.mutex.Unlock()

	return
}

func (logger *Logger) Close() (err error) {
	if logger == nil {
		err = fmt.Errorf("logger is nil")
		return
	}

	var str string
	str = time.Now().Format(DeafultTimeFormat)
	logger.writer.Write([]byte(str + " ~~~ LOGGER EXIT\n"))

	if logger.file != nil {
		logger.file.Close()
	}

	return
}

//// Log Error
// print error to console when err != nil for debug
func PrintErr(tag string, err error, items ...interface{}) {
	if err == nil {
		return
	}

	fmt.Printf("\n>>> %s: %v\n", tag, err)

	for i := range items {
		fmt.Printf("  %#v\n", items[i])
	}

	fmt.Println("<<<")
}

// log error, print error even it's nil
func LogErr(tag string, err error) {
	var at, prefix string
	at = time.Now().Format("2006-01-02T15:04:05-0700")

	if err == nil {
		prefix = ">>>"
	} else {
		prefix = "!!!"
	}

	fmt.Printf("\n[%s] %s %s: %v\n", at, prefix, tag, err)
}
