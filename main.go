package main

import (
	"log"
	"net/http"
	"os"

	"github.com/KentoSakurai-UEC/fib_api/fiblogic"
	handlerV1 "github.com/KentoSakurai-UEC/fib_api/handler/handler_v1"
)

func main() {
	twoStepFibCalc := fiblogic.NewTwoStepFibCalc()

	v1handler := handlerV1.NewHandler(twoStepFibCalc)

	http.HandleFunc("/fib", v1handler.FibHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on: %s\n", port)
}
