package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// reference: https://stackoverflow.com/questions/29995741/how-can-i-source-a-shell-script-using-go
func LoadBashrc(fp string) (err error) {
	var (
		start   bool
		bs      []byte
		kv      []string
		cmd     *exec.Cmd
		scanner *bufio.Scanner
	)

	cmd = exec.Command("bash", "-c",
		fmt.Sprintf(`set -e; source "%s"; echo "<<<ENVIRONMENT>>>"; env`, fp))

	if bs, err = cmd.CombinedOutput(); err != nil {
		return
	}

	scanner = bufio.NewScanner(bytes.NewReader(bs))
	for scanner.Scan() {
		if start {
			if kv = strings.SplitN(scanner.Text(), "=", 2); len(kv) == 2 {
				os.Setenv(kv[0], kv[1])
			}
			continue
		}
		if scanner.Text() == "<<<ENVIRONMENT>>>" {
			start = true
		}
	}

	return
}
