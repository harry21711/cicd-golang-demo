package main

import (
	"fmt"
	"os"
)

var JenkinsBuildInfo string

var BuildURL = os.Getenv("BUILD_URL")
var BuildNo = os.Getenv("BUILD_NUMBER")
// var BuildURL = "http://localhost:8080/job/Clone%20Simple/12/"
// var BuildNo = "12"

func GetJenkinsInfo() string {
	if BuildNo != "" {
		JenkinsBuildInfo = JenkinsBuildInfo + "\nBuild No : " + BuildNo
	}
	if BuildURL != "" {
		JenkinsBuildInfo = JenkinsBuildInfo + "\nBuild URL : " + BuildURL
	}
	return JenkinsBuildInfo
}

func main() {
	//fmt.Println("start")
	//for _, env := range os.Environ() {
		//fmt.Println(env)
	//}
	//fmt.Println("end")
	comment := "Update configs"
	comment = comment + GetJenkinsInfo()
	
	fmt.Println(comment)
}

