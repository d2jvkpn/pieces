package misc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/pprof"
	"runtime"
	"time"
)

/*
web browser address
  http://localhost:5060/debug/pprof/

get profiles and view in browser
  $ go tool pprof -http=:8080 http://localhost:5060/debug/pprof/allocs?seconds=30
  $ go tool pprof http://localhost:5060/debug/pprof/block?seconds=30
  $ go tool pprof http://localhost:5060/debug/pprof/goroutine?seconds=30
  $ go tool pprof http://localhost:5060/debug/pprof/heap?seconds=30
  $ go tool pprof http://localhost:5060/debug/pprof/mutex?seconds=30
  $ go tool pprof http://localhost:5060/debug/pprof/profile?seconds=30
  $ go tool pprof http://localhost:5060/debug/pprof/threadcreate?seconds=30

download profile file and convert to svg image
  $ wget -O profile.out localhost:5060/debug/pprof/profile?seconds=30
  $ go tool pprof -svg profile.out > profile.svg

get pprof in 30 seconds and save to svg image
  $ go tool pprof -svg http://localhost:5060/debug/pprof/allocs?seconds=30 > allocs.svg

get trace in 5 seconds
  $ wget -O trace.out http://localhost:5060/debug/pprof/trace?seconds=5
  $ go tool trace trace.out

get cmdline and symbol binary data
  $ wget -O cmdline.out http://localhost:5060/debug/pprof/cmdline
  $ wget -O symbol.out http://localhost:5060/debug/pprof/symbol
*/
type Pprof struct {
	addr   string
	Server *http.Server
	status string
	err    error
}

// create new Pprof and run server
func NewPprof(addr string) (pp *Pprof) {
	mux := http.NewServeMux()

	mux.HandleFunc("/debug/healthy", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte{})
	})

	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)

	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	mux.HandleFunc("/debug/runtime/status", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Content-Type", "application/json; charset=utf-8")

		memStats := new(runtime.MemStats)
		runtime.ReadMemStats(memStats)
		num := runtime.NumGoroutine()

		json.NewEncoder(res).Encode(map[string]interface{}{
			"numGoroutine": num,
			"memStats":     memStats,
		})
	})

	pp = &Pprof{addr: addr, status: "running"}
	pp.Server = &http.Server{
		Addr:        addr,
		Handler:     mux,
		ReadTimeout: 5 * time.Second,
		// WriteTimeout:   100 * time.Second,
		MaxHeaderBytes: 1 << 23, // 8M
	}

	return
}

func (pp *Pprof) Run() (err error) {
	pp.status = "running"

	pp.err = pp.Server.ListenAndServe()

	if pp.err != http.ErrServerClosed {
		pp.status = "failed"
	} else {
		pp.status = "shutdown"
	}

	return pp.err
}

// print add, status and err in json format
func (pp *Pprof) String() string {
	if pp.err == nil {
		return fmt.Sprintf(`{"addr": %q, "status": %q}`, pp.addr, pp.status)
	}

	return fmt.Sprintf(
		`{"addr": %q, "status": %q, "err": %q}`,
		pp.addr, pp.status, pp.err.Error(),
	)
}

// get err field
func (pp *Pprof) Err() (err error) {
	return pp.err
}

// showdown server
func (pp *Pprof) Shutdown() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if err = pp.Server.Shutdown(ctx); err != nil {
		pp.status = "failed"
	} else {
		pp.status = "shutdown"
	}
	cancel()
	return err
}
