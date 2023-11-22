package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func main() {

	log.Printf
	log.Printf("Main Start")
	buildNumber := os.Getenv("BUILD_NUMBER")
	log.Printf("buildNumber : %s", buildNumber)
	fmt.Println("Main End")

	// r := mux.NewRouter()
	// r.HandleFunc("/", Output)
	// http.Handle("/", r)

	// log.Fatal(http.ListenAndServe(":9000", r))
}

func Output(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Random no : %v", randomNumbers())
}

func randomNumbers() int {
	return rand.Intn(1000)
}

func add(a, b int) int {
	return a + b
}
