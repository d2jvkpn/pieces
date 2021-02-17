package misc

import (
	"fmt"
	"os"
	"path/filepath"
)

// get root path of project by recursively match go.mod in parent directory
func ProjectDir() (p string, err error) {
	var (
		x, tmp string
		ms     []string
	)
	if x, err = os.Getwd(); err != nil {
		return "", err
	}

	for {
		if ms, err = filepath.Glob(filepath.Join(x, "go\\.mod")); err != nil {
			return "", nil
		}
		if len(ms) > 0 {
			return x, nil
		}

		if tmp = filepath.Dir(x); x == tmp {
			break
		}
		x = tmp
	}

	return "", fmt.Errorf("not found project dir")
}
