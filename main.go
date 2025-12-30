package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/tylerBrittain42/blog/pkg/validator"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	cfg := config{templateDir: os.Getenv("DIR")}

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("GET /article/{name}", cfg.articleHandler)

	log.Printf("Serving on port %s\n", port)
	log.Fatal(server.ListenAndServe())

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "template/index.html")
}

func (cfg *config) articleHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	isSanitized, err := validator.IsAlphaNumeric(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error: %v\n", err)))
		return
	}
	if isSanitized == false {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Invalid characters in name"))
		return
	}

	canAccess, err := helper.IsAccessable(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error: %v\n", err)))
		return
	}
	if canAccess == false {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Unable to access article"))
		return
	}

	http.ServeFile(w, r, "template/index.html")


type config struct {
	templateDir string
}
