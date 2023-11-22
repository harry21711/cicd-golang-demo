package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("start")
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
	fmt.Println("end")
	fmt.Println(os.Getenv("BU"))
	fmt.Println("final end")
}

