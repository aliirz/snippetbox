package main

import (
	"log"
	"net/http"
)

// home handler function
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Go"))
}

func main() {
	//use http.NewServerMux() to handle home on "/"
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting server on :1337")
	err := http.ListenAndServe(":1337", mux)
	log.Fatal(err)
}
