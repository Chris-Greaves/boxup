package main

import (
	"log"
	"os"
	"net/http"
)

var (
	Logger = log.New(os.Stderr, "BoxUp: ", log.Lshortfile|log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
)

func main() {
	router := NewRouter()
	Logger.Println("Service starting on port 5950")
	Logger.Fatal(http.ListenAndServe(":5950", router))
}