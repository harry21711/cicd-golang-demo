package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func main() {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}

