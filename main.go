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

	comment := "Update configs"
	buildNo := os.Getenv("BN")
	buildURL := os.Getenv("BU")
	
	if buildNo != "" {
		comment = comment + "\nbuild no : " + buildNo
	}

	if buildURL != "" {
		comment = comment + "\nbuild job url : " + buildURL
	}

	fmt.Println(comment)
}

