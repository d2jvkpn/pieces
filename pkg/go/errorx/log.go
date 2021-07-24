package errorx

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Logger struct {
	*log.Logger
	printCaller bool
	setStdout   bool
	prefix      string
	format      string
	current     string
	tag         string
	file        *os.File
	ch          chan bool
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

func NewLogger2(prefix, format string) (lg *Logger, err error) {
	if lg, err = NewLogger(prefix, format); err != nil {
		return nil, err
	}

	lg.printCaller = true
	return lg, err
}

func (lg *Logger) Close() (err error) {
	<-lg.ch //!! important
	close(lg.ch)
	return lg.file.Close()
}

func (lg *Logger) SetStdout() {
	lg.setStdout = true
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
	RFC3339ms := "2006-01-02T15:04:05.000Z07:00 " // time.RFC3339
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

	buf := bytes.NewBufferString(now.Format(RFC3339ms))

	if lg.printCaller {
		buf.WriteString(CallInfo(3))
		buf.WriteString(" ")
	}
	buf.Write(bts)

	var wr io.Writer
	if lg.setStdout {
		wr = io.MultiWriter(os.Stdout, lg.file)
	} else {
		wr = io.MultiWriter(lg.file)
	}

	n, err = wr.Write(buf.Bytes())
	buf.Reset()
	return
}

func (lg *Logger) Info(format string, a ...interface{}) (n int, err error) {
	buf := bytes.NewBufferString("[INFO] ")
	buf.Write(bytes.TrimSpace([]byte(format)))
	buf.WriteByte('\n')

	if len(a) > 0 {
		n, err = lg.Write([]byte(fmt.Sprintf(buf.String(), a...)))
	} else {
		n, err = lg.Write(buf.Bytes())
	}
	buf.Reset()
	return
}

func (lg *Logger) InfoBytes(bts []byte) (n int, err error) {
	buf := bytes.NewBuffer([]byte("[INFO] "))
	buf.Write(bts)
	n, err = lg.Write(buf.Bytes())
	buf.Reset()
	return
}

func (lg *Logger) Warn(format string, a ...interface{}) (n int, err error) {
	buf := bytes.NewBufferString("[WARN] ")
	buf.Write(bytes.TrimSpace([]byte(format)))
	buf.WriteByte('\n')

	if len(a) > 0 {
		n, err = lg.Write([]byte(fmt.Sprintf(buf.String(), a...)))
	} else {
		n, err = lg.Write(buf.Bytes())
	}
	buf.Reset()
	return
}

func (lg *Logger) Error(format string, a ...interface{}) (n int, err error) {
	format = strings.TrimSpace("[ERROR] "+format) + "\n"

	if len(a) > 0 {
		return lg.Write([]byte(fmt.Sprintf(format, a...)))
	} else {
		return lg.Write([]byte(format))
	}
}

func (lg *Logger) Panic(format string, a ...interface{}) (n int, err error) {
	format = strings.TrimSpace("[PANIC] "+format) + "\n"

	if len(a) > 0 {
		return lg.Write([]byte(fmt.Sprintf(format, a...)))
	} else {
		return lg.Write([]byte(format))
	}
}
