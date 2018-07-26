package main

import (
	"log"
	"net/http"

	handler "github.com/chilledblooded/gophercises/Exercise_15/Handlers"
	recovery "github.com/chilledblooded/gophercises/Exercise_15/Middleware"
)

func main() {

	log.Fatal(http.ListenAndServe(":8888", recovery.RecoveryMid(handler.Handler())))
}
