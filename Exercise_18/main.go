package main

import (
	"fmt"

	"github.com/chilledblooded/gophercises/Exercise_18/primitive"
)

func main() {
	_, err := primitive.RunPrimitive("./img/gopher.jpeg", "./img/out.jpeg", 50)
	if err != nil {
		fmt.Printf("Failed to run primitive command %v", err)
	}
}
