package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	var (
		addr string
		secs int
	)

	flag.StringVar(&addr, "addr", "http://127.0.0.1:1030", "pprof server address")
	flag.IntVar(&secs, "secs", 30, "pprof request seconds")
	flag.Parse()

	paths := []string{
		"/debug/pprof/allocs",
		"/debug/pprof/block",
		"/debug/pprof/goroutine",
		"/debug/pprof/heap",
		"/debug/pprof/mutex",
		"/debug/pprof/profile",
		"/debug/pprof/threadcreate",
		"/debug/pprof/trace",
	}

	client := new(http.Client)
	wg := new(sync.WaitGroup)
	now := time.Now()

	dir := filepath.Join(
		"logs",
		fmt.Sprintf("%d_%s_pprof", now.Unix(), now.Format("2006-01-02T15-04-05")),
	)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalln(err)
	}

	download := NewDownload(client, secs, dir, wg)

	for _, p := range paths {
		wg.Add(1)

		go func(p string) {
			if err := download(addr + p); err != nil {
				log.Printf("%v\n", err)
			}
		}(p)
	}

	wg.Wait()
	log.Println("done")
}

func NewDownload(client *http.Client, secs int, dir string, wg *sync.WaitGroup) func(string) error {
	return func(p string) (err error) {
		var (
			resp *http.Response
			file *os.File
		)
		defer wg.Done()

		link := fmt.Sprintf("%s?seconds=%d", p, secs)
		log.Printf("download %s\n", link)

		if resp, err = client.Get(link); err != nil {
			return err
		}
		defer resp.Body.Close()

		if file, err = os.Create(filepath.Join(dir, filepath.Base(p)+".out")); err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(file, resp.Body)
		return err
	}
}
