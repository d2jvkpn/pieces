package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	cmd := exec.Command(
		"/bin/bash", "-c",
		`n=0; while true; do n=$((n+1)); echo "    $n"; sleep 1; done`,
	)

	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

	fmt.Println(">>> run command")
	cmd.Start()

	time.AfterFunc(10*time.Second, func() {
		fmt.Println("    kill process")
		cmd.Process.Kill()
	})
	cmd.Wait()

	fmt.Println("exit program.")
	os.Exit(0)
}
