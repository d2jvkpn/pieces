package misc

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func NewError2(err error) error {
	if err == nil {
		return nil
	}

	fn, file, line, _ := runtime.Caller(1)

	return fmt.Errorf(
		"%s (%s [%d]): %w", runtime.FuncForPC(fn).Name(),
		filepath.Base(file), line, err,
	)

}
