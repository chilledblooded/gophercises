package main

import (
	"log"
	"net/http"

	handler "github.com/chilledblooded/gophercises/Exercise_18/handlers"
)

var listenAndServeFunc = http.ListenAndServe

func main() {
	log.Fatal(listenAndServeFunc(":8888", handler.GetMux()))
}
