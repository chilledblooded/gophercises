package main

import (
	"log"
	"net/http"

	handler "github.com/chilledblooded/gophercises/Exercise_15/Handlers"
	recovery "github.com/chilledblooded/gophercises/Exercise_15/Middleware"
)

var listenAndServeFunc = http.ListenAndServe

func main() {

	log.Fatal(listenAndServeFunc(":8888", recovery.RecoveryMid(handler.Handler())))
}
