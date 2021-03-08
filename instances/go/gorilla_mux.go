package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fileName := "pages/" + vars["id"] + ".html"
	fmt.Println(">>> serving", fileName)
	http.ServeFile(w, r, fileName)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("src/pages/{id:[0-9]+}", pageHandler)

	// http.Handle("/", router)
	server := http.Server{
		Addr:    "localhost:8010",
		Handler: router,
	}

	fmt.Println("Please visit :8010/pages/{number}")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
