package errorx

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Logger struct {
	*log.Logger
	prefix  string
	format  string
	current string
	file    *os.File
	ch      chan bool
}

func NewLogger(prefix, format string) (lg *Logger, err error) {
	if !strings.HasSuffix(prefix, ".") {
		prefix += "."
	}

	lg = &Logger{
		prefix:  prefix,
		format:  format,
		current: time.Now().Format(format),
		ch:      make(chan bool, 1),
	}

	if err = os.MkdirAll(filepath.Dir(lg.prefix), 0755); err != nil {
		return nil, err
	}

	lg.file, err = os.OpenFile(
		fmt.Sprintf("%s%s.log", lg.prefix, lg.current),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644,
	)
	if err != nil {
		return nil, err
	}

	lg.Logger = new(log.Logger)
	lg.Logger.SetOutput(lg)
	lg.ch <- true

	return lg, nil
}

func (lg *Logger) Close() (err error) {
	<-lg.ch //!!
	close(lg.ch)
	return lg.file.Close()
}

func (lg *Logger) rotate(now time.Time) (err error) {
	target := now.Format(lg.format)
	if target == lg.current {
		return nil
	}

	if err = lg.file.Close(); err != nil {
		return err
	}

	lg.file, err = os.OpenFile(
		fmt.Sprintf("%s%s.log", lg.prefix, lg.current),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644,
	)

	return err
}

func (lg *Logger) Write(bts []byte) (n int, err error) {
	var RFC3339ms = "2006-01-02T15:04:05.000Z07:00" // time.RFC3339
	now := time.Now()

	if ok := <-lg.ch; !ok {
		return 0, fmt.Errorf("logger was closed")
	}
	defer func() {
		lg.ch <- true
	}()

	if err = lg.rotate(now); err != nil {
		return 0, err
	}

	n, err = fmt.Fprintf(
		lg.file, "%s %s\n",
		now.Format(RFC3339ms), strings.TrimSpace(string(bts)),
	)

	return n, err
}
