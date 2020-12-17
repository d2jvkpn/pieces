package misc

import (
	"context"
	"fmt"
	"net/http"
	"net/http/pprof"
	"time"
)

/*

  web browser address
    http://localhost:8080/debug/pprof/


  get profiles and view in browser
  ```bash
  $ go tool pprof http://localhost:8080/debug/pprof/allocs?seconds=30
  $ go tool pprof http://localhost:8080/debug/pprof/block?seconds=30
  $ go tool pprof http://localhost:8080/debug/pprof/goroutine?seconds=30
  $ go tool pprof http://localhost:8080/debug/pprof/heap?seconds=30
  $ go tool pprof http://localhost:8080/debug/pprof/mutex?seconds=30
  $ go tool pprof http://localhost:8080/debug/pprof/profile?seconds=30
  $ go tool pprof http://localhost:8080/debug/pprof/threadcreate?seconds=30
  ```

  download profile file and convert to svg image
  ```bash
  $ wget -O profile.out localhost:8080/debug/pprof/profile?seconds=30
  $ go tool pprof  -svg profile.out > profile.svg
  ```

  get pprof in 3o0 seconds svg image
  $ go tool pprof -svg http://localhost:8080/debug/pprof/allocs?seconds=30 > allocs.svg

  get trace in 5 seconds
  ```bash
  $ wget -O trace.out http://localhost:8080/debug/pprof/trace?seconds=5
  $ go tool trace trace.out
  ```

  get cmdline and symbo binary data
  ```bash
  $ wget -O cmdline.out http://localhost:8080/debug/pprof/cmdline
  $ wget -O symbol.out http://localhost:8080/debug/pprof/symbol
  ```
*/
type Pprof struct {
	addr   string
	server *http.Server
	status string
	err    error
}

// create new Pprof and run server
func NewPprof(addr string) (pp *Pprof) {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	pp = &Pprof{
		addr:   addr,
		status: "running",
	}

	pp.server = &http.Server{
		Addr:        addr,
		Handler:     mux,
		ReadTimeout: 5 * time.Second,
		// WriteTimeout:   100 * time.Second,
		MaxHeaderBytes: 1 << 23, // 8M
	}

	go func() {
		pp.err = pp.server.ListenAndServe()

		if pp.err != http.ErrServerClosed {
			pp.status = "failed"
		} else {
			pp.status = "shutdown"
		}
	}()
	return
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
	if err = pp.server.Shutdown(ctx); err != nil {
		pp.status = "failed"
	} else {
		pp.status = "shutdown"
	}
	cancel()
	return err
}
