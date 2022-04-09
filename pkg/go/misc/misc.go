package misc

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	// "strconv"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"golang.org/x/exp/constraints"
)

///
var (
	Rand            *rand.Rand
	_base64Encoding *base64.Encoding = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_")
)

const (
	RFC3339ms = "2006-01-02T15:04:05.000Z07:00"
)

func init() {
	Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	// rand.Seed(time.Now().UnixNano()); rand.Fn()
}

func quoteString(str string) (out string) {
	if str == "" {
		return ""
	}
	out = fmt.Sprintf("%q", str)
	out = out[1 : len(out)-1] // remove quote
	return
}

func NewRand() (rd *rand.Rand) {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func PrintJSON(d interface{}) (err error) {
	coder := json.NewEncoder(os.Stdout)
	coder.SetIndent("", "  ")

	return coder.Encode(d)
}

///
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

///
func ListenOSIntr(do func(), errch chan<- error, sgs ...os.Signal) {
	// linux support syscall.SIGUSR2
	quit := make(chan os.Signal, 1)

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

///
func StrsIndex(ss []string, s string) (i int) {
	for i = range ss {
		if ss[i] == s {
			return i
		}
	}

	return -1
}

///
func SetLogRFC3339() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
}

type logWriter struct{}

func (writer *logWriter) Write(bts []byte) (int, error) {
	// time.RFC3339
	return fmt.Print(time.Now().Format(RFC3339ms) + " " + string(bts))
}

///
func QuoteString(str string) string {
	if str == "" {
		return ""
	}
	return fmt.Sprintf("%q", str)
}

func FileSize2Str(n int64) string {
	switch {
	case n <= 0:
		return "0"
	case n < 1<<10:
		return fmt.Sprintf("%dB", n)
	case n >= 1<<10 && n < 1<<20:
		return fmt.Sprintf("%dK", n>>10)
	case n >= 1<<20 && n < 1<<30:
		return fmt.Sprintf("%dM", n>>20)
	default:
		return fmt.Sprintf("%dG", n>>30)
	}
}

func VectorIndex[T constraints.Ordered](list []T, v T) int {
	for i := range list {
		if list[i] == v {
			return i
		}
	}

	return -1
}

func EqualVector[T constraints.Ordered](arr1, arr2 []T) (ok bool) {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

func UniqVector[T constraints.Ordered](arr []T) (list []T) {
	n := len(arr)
	list = make([]T, 0, n)

	if len(arr) == 0 {
		return list
	}

	mp := make(map[T]bool, n)
	for _, v := range arr {
		if !mp[v] {
			list = append(list, v)
			mp[v] = true
		}
	}

	return list
}

// replace +/ with -_
func Base64Encode(src []byte) string {
	return _base64Encoding.EncodeToString(src)
}

// replace +/ with -_
func Base64Decode(src string) ([]byte, error) {
	return _base64Encoding.DecodeString(src)
}

func FileSaveName(p string) (out string, err error) {
	var (
		i         int
		base, ext string
	)

	ext = filepath.Ext(p)
	base = p[0:(len(p) - len(ext))]
	i, out = 1, p
	for {
		// fmt.Println(i, out)
		if _, err = os.Stat(out); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				return out, nil
			}
			return "", err
		}
		i++
		out = fmt.Sprintf("%s-%d%s", base, i, ext)
	}
}
