package main

import (
	"log"
	"net/http"

	handler "github.com/chilledblooded/gophercises/Exercise_18/handlers"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./img/"))
	mux.Handle("/img/", http.StripPrefix("/img", fs))
	mux.HandleFunc("/", handler.Home)
	mux.HandleFunc("/upload", handler.Upload)
	mux.HandleFunc("/modify/", handler.Modify)
	log.Fatal(http.ListenAndServe(":8888", mux))
}
