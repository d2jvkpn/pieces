package misc

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHttpError_t1(t *testing.T) {
	var httpErr = NewHttpError

	err := httpErr(
		fmt.Errorf("something is wrong"), "request failed",
		http.StatusInternalServerError, 1,
	)

	PrintJSON(err)
}

func TestHttpError_t2(t *testing.T) {
	var httpErr = NewHttpError

	err := httpErr(
		fmt.Errorf("something is wrong"), "request failed",
		http.StatusInternalServerError, 1,
	)

	HttpErrorResetCode(err, 10)

	PrintJSON(err)
}

func TestHttpError_t3(t *testing.T) {
	var httpErr = NewHttpError

	err := httpErr(
		fmt.Errorf("invlaid parameter found"), "bad request",
		http.StatusBadRequest, -1,
	)

	HttpErrorResetCode(err, 10, -5)

	PrintJSON(err)
}

func TestHttpError_t4(t *testing.T) {
	var httpErr = NewHttpError

	err := httpErr(
		fmt.Errorf("invlaid parameter found"), "bad request",
		http.StatusBadRequest, -1,
	)

	fmt.Println(HttpErrorExtract(err))
}
