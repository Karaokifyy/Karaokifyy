package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello-world", helloWorld)

	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    ":http",
	}

	log.Fatal(srv.ListenAndServe())
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello")
	var data = struct {
		Title string `json:"title"`
	}{
		Title: "See if this works THIS IS A TEST",
	}

	jsonBytes, err := StructToJSON(data)
	if err != nil {
		fmt.Print(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
