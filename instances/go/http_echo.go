package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"

	uuid "github.com/satori/go.uuid"
)

func Hi(w http.ResponseWriter, r *http.Request) {
	var (
		keys   []string
		msg    string
		path   string
		err    error
		cookie *http.Cookie
	)

	if err = r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("%s\n", err), 400)
		return
	}

	///
	fmt.Fprintf(w,
		">>> Hello, %s\n    RemoteAddr: %q\n    Host: %q\n",
		time.Now().Format("2006-01-02T15:04:05.000-07:00"), r.RemoteAddr, r.Host,
	) // time.Now().Format(time.RFC3339)

	fmt.Fprintf(w,
		"    Method: %q\n    Request: %q\n    RawQuery: %q\n\n",
		r.Method, r.URL.Path, r.URL.RawQuery,
	) // r.RequestURI

	///
	path, msg = "Hi", ">>> Cookie:\n    cookie named %s had been set to %q.\n\n"
	if cookie, err = r.Cookie(path); err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{Name: path, Value: id.String()}

		http.SetCookie(w, cookie)
		msg = ">>> Cookie:\n    no cookie named %s, and it was set to %q.\n\n"

	}
	fmt.Fprintf(w, msg, path, cookie.Value)

	/// Data
	fmt.Fprintf(w, ">>> Data:\n")
	for k, _ := range r.Form {
		fmt.Fprintf(w, "      %s = `%#v`\n", k, r.Form[k])
	}
	fmt.Fprintln(w, "\n")

	/// Header
	fmt.Fprintf(w, ">>> Header:\n")
	for k, _ := range r.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Fprintf(w, "      %s = `%#v`\n", k, r.Header[k])
	}
	fmt.Fprintln(w, "")

	return
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hi)

	server := http.Server{
		Addr:    "localhost:8010",
		Handler: mux,
	}

	fmt.Println("Please visit :8010")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
