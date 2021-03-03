package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

var DeafultTimeFormat = "[2006-01-02T15:04:05-0700]"

type Argx struct {
	Start
	End
	*Logger
}

type Start struct {
	Main     string
	WorkPath string
	User     string
	PID      int
	Created  time.Time
	Args     []string
}

type End struct {
	Code    int
	Message string
	EndAt   time.Time
	Elapsed string
}

// save command arguments to data struct Argx
func NewArgx() (argx *Argx, err error) {
	var u *user.User

	argx = new(Argx)
	argx.Created, argx.PID = time.Now(), os.Getpid()
	argx.Args = make([]string, len(os.Args))
	copy(argx.Args, os.Args)

	// println(argx.Args)
	argx.Main, _ = filepath.Abs(argx.Args[0])
	if argx.WorkPath, err = os.Getwd(); err != nil {
		return
	}

	if u, err = user.Current(); err != nil {
		return
	}

	argx.User = fmt.Sprintf("{\"uid\":%q, \"gid\":%q, \"username\":%q}",
		u.Uid, u.Gid, u.Username)

	argx.Logger = NewLogger()

	return
}

// save end information to argx.End
func (argx *Argx) Done(code int, msg string) {
	argx.Code = code
	argx.Message = msg
	argx.EndAt = time.Now()
	argx.Elapsed = argx.EndAt.Sub(argx.Created).String()
	argx.Logger.Close()

	return
}
