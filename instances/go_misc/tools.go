package rover

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/tabwriter"
)

//// Save bytes and string to file
// save []byte to file (and create parent directories)
func Bts2File(bts []byte, out string) (err error) {
	var file *os.File

	if err = os.MkdirAll(path.Dir(out), 0755); err != nil {
		return
	}

	if file, err = os.Create(out); err != nil {
		return
	}
	defer file.Close()

	_, err = file.Write(bts)
	return
}

func Str2File(str, out string) (err error) {
	err = Bts2File([]byte(str), out)
	return
}

//// Path
// print slice of string slice in align text
func Columns(array [][]string, prefixSpaces, padding int) {
	var (
		i      int
		writer *tabwriter.Writer
	)

	writer = tabwriter.NewWriter(os.Stdout, 4, 0, padding, ' ',
		tabwriter.StripEscape)

	for i = range array {
		fmt.Fprintln(writer, strings.Repeat(" ", prefixSpaces)+
			strings.Join(array[i], "\t"))
	}

	writer.Flush()
}

// return file parent directory, basename and extend name
func DecomposePath(p string) (dir, base, ext string, err error) {
	// os.IsNotExist
	var abs string
	if _, err = os.Stat(p); err != nil {
		return
	}

	base, ext = filepath.Base(p), filepath.Ext(p)
	base = strings.TrimSuffix(base, ext)
	abs, _ = filepath.Abs(p)
	dir = filepath.Dir(abs)
	return
}

// return file realpath
func RealPath(p string) (realpath string, err error) {
	if _, err = os.Stat(p); err != nil {
		return
	}

	if realpath, err = filepath.EvalSymlinks(p); err != nil {
		return
	}

	realpath, _ = filepath.Abs(realpath)

	return
}

//// MD5
var (
	md5RE      = "^[a-zA-Z0-9]{32}$"
	md5UpperRE = "^[A-Z0-9]{32}$"
	md5LowerRE = "^[a-z0-9]{32}$"
)

func IsMD5(str string) (err error) {
	var match bool
	if len(str) != 32 {
		err = fmt.Errorf("string lenght isn't 32")
		return
	}

	match, _ = regexp.Match(md5RE, []byte(str))
	if !match {
		err = fmt.Errorf("not a valid md5sum")
		return
	}

	return
}

func IsMD5Upper(str string) (err error) {
	var match bool
	if len(str) != 32 {
		err = fmt.Errorf("string lenght isn't 32")
		return
	}

	match, _ = regexp.Match(md5UpperRE, []byte(str))
	if !match {
		err = fmt.Errorf("not a valid upper case md5sum")
		return
	}

	return
}

func IsMD5Lower(str string) (err error) {
	var match bool
	if len(str) != 32 {
		err = fmt.Errorf("string lenght isn't 32")
		return
	}

	match, _ = regexp.Match(md5LowerRE, []byte(str))
	if !match {
		err = fmt.Errorf("not a valid lower case md5sum")
		return
	}

	return
}

//// Index
func IndexString(slice []string, value string) (p int) {
	for p = range slice {
		if slice[p] == value {
			return
		}
	}

	p = -1
	return
}

func IndexInt(slice []int, value int) (p int) {
	for p = range slice {
		if slice[p] == value {
			return
		}
	}

	p = -1
	return
}

func IndexInt64(slice []int64, value int64) (p int) {
	for p = range slice {
		if slice[p] == value {
			return
		}
	}

	p = -1
	return
}

func BindError(e1, e2 error) (err error) {
	switch {
	case e1 == nil && e2 == nil:
	case e1 != nil && e2 != nil:
		err = fmt.Errorf("%v: %w", e1, e2)
	case e1 != nil && e2 != nil:
		err = e1
	case e1 == nil && e2 != nil:
		err = e2
	}

	return
}
