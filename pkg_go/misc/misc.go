package misc

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

//
var (
	Rand *rand.Rand
)

const (
	RFC3339ms = "2006-01-02T15:04:05.000Z07:00"
)

func init() {
	Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	// rand.Seed(time.Now().UnixNano()); rand.Fn()
}

func NewRand() (rd *rand.Rand) {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func PrintJSON(d interface{}) (err error) {
	coder := json.NewEncoder(os.Stdout)
	coder.SetIndent("", "  ")

	return coder.Encode(d)
}

//
func AddHttpsScheme(a string) string {
	a = strings.TrimSpace(a)

	switch {
	case strings.HasPrefix(a, "https://") || a == "":
		return a

	case strings.HasPrefix(a, "//"):
		return "https:" + a

	case strings.HasPrefix(a, "http:"):
		return "https:" + a[5:]
	}

	return "https://" + a
}

//
func ListenOSIntr(do func(), errch chan<- error, sgs ...os.Signal) {
	// linux support syscall.SIGUSR2
	quit := make(chan os.Signal)

	if len(sgs) == 0 {
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	} else {
		signal.Notify(quit, sgs...)
	}
	<-quit

	if do != nil {
		do()
	}
	if errch != nil { //! errch == nil means doesn't need send an nil to channel
		errch <- nil
	}

	return
}

//
func StrsIndex(ss []string, s string) (i int) {
	for i = range ss {
		if ss[i] == s {
			return i
		}
	}

	return -1
}

//
func SetLogRFC3339() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
}

type logWriter struct{}

func (writer *logWriter) Write(bts []byte) (int, error) {
	// time.RFC3339
	return fmt.Print(time.Now().Format(RFC3339ms) + "  " + string(bts))
}
