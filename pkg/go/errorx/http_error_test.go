package errorx

import (
	"fmt"
	"net/http"
	"testing"
)

var (
	httpErr = NewHttpError  // func(error, string, int, int)
	resBad  = ResBadRequest // func(http.ResponseWriter, int, error)
	resOk   = ResOk         // func(http.ResponseWriter)
	// resJSON = ResJSON
)

func TestHttpError_t1(t *testing.T) {
	var httpErr = NewHttpError

	err := httpErr(
		fmt.Errorf("something is wrong"), "request failed",
		http.StatusInternalServerError, 1,
	)

	fmt.Println(err)
}

func TestHttpError_t2(t *testing.T) {
	var httpErr = NewHttpError

	err := httpErr(
		fmt.Errorf("something is wrong"), "request failed",
		http.StatusInternalServerError, 1,
	)

	HttpErrorResetCode(err, 10)

	fmt.Println(err)
}

func TestHttpError_t3(t *testing.T) {
	var httpErr = NewHttpError

	err := httpErr(
		fmt.Errorf("invlaid parameter found"), "bad request",
		http.StatusBadRequest, -1,
	)

	HttpErrorResetCode(err, 10, -5)

	fmt.Println(err)
}

func TestHttpError_t4(t *testing.T) {
	var httpErr = NewHttpError

	err := httpErr(
		fmt.Errorf("invlaid parameter found"), "bad request",
		http.StatusBadRequest, -1,
	)

	fmt.Println(HttpErrorExtract(err))
}
