package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	buildNumber := os.Getenv("BUILD_NUMBER")
	log.Printf("drift type: %s", buildNumber)
	fmt.Println("buildNumber : ", buildNumber)

	r := mux.NewRouter()
	r.HandleFunc("/", Output)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":9000", r))
}

func Output(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Random no : %v", randomNumbers())
}

func randomNumbers() int {
	return rand.Intn(1000)
}
