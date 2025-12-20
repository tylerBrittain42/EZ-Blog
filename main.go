package main

import (
	"log"
	"net/http"

	"github.com/tylerBrittain42/blog/pkg/helper"
)

func main() {
	port := "8080"
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("GET /article/{name}", articleHandler)

	log.Printf("Serving on port %s\n", port)
	log.Fatal(server.ListenAndServe())

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from a byte string"))
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	isValidName, err := helper.IsAlphaNumeric(name)
	if err != nil || isValidName != true {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("error or does not meet validation"))
		return
	}
	// w.Write([]byte("passes validation for " + name))
	http.ServeFile(w, r, "template/index.html")

}
