package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var (
		wd, localPath   string
		prefix, address string
		err             error
	)

	///
	if wd, err = os.Getwd(); err != nil {
		fatalf("os.Getwd(): %v", err)
	}

	flag.StringVar(&localPath, "path", wd, "local directory path")
	flag.StringVar(&prefix, "prefix", "/", "http path prefix")
	flag.StringVar(&address, "address", ":8000", "serve address")
	flag.Parse()

	if localPath, err = filepath.Abs(localPath); err != nil {
		fatalf("filepath.Abs(?): %v", err)
	}

	///
	mux := http.NewServeMux()
	mux.Handle(prefix, http.StripPrefix(
		"/"+strings.Trim(prefix, "/"),
		http.FileServer(http.Dir(localPath)),
	))

	server := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	///
	log.Printf("Serving %q", localPath)
	log.Printf("HTTP listening on %q\n", address)
	if err = server.ListenAndServe(); err == http.ErrServerClosed {
		log.Println("exit program")
		os.Exit(0)
	}
	fatalf("http ListenAndServe(): %v", err)
}

func fatalf(format string, v ...interface{}) {
	log.Fatalf(strings.TrimRight(format, "\n ")+"\n", v...)
}
