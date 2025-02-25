package main

import (
	"log"
	"net/http"
	"os"

	handlerV1 "github.com/KentoSakurai-UEC/fib_api/handler/handler_v1"
	"github.com/KentoSakurai-UEC/fib_api/logic"
)

func main() {
	twoStepFibCalc := logic.NewTwoStepFibCalc()
	businessLogic := logic.NewBusinessLogic(twoStepFibCalc)

	v1handler := handlerV1.NewHandler(businessLogic)
	http.HandleFunc("/fib", v1handler.FibHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on: %s\n", port)
}
