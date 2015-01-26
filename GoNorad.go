package main

import (
	"fmt"
	"net/http/cookiejar"
	"os"
)

type DataRequest struct {
	result []string
}

func main() {

	var res string
	var e error
	var np2 Neptune
	config := GetConfig()
	cookieJar, e := cookiejar.New(nil)

	res, e = np2.Login(config["username"], config["password"], cookieJar)
	ErrorExit(e)

	fmt.Printf("%v - %v", cookieJar, res)
	fmt.Println()

	res, e = np2.GetData(config["gameNumber"], cookieJar)
	ErrorExit(e)

	fmt.Printf("%v - %v", cookieJar, res)
	fmt.Println()

}

func ErrorExit(e error) {
	if e != nil {
		fmt.Println("Error:", e)
		os.Exit(1)
	}
}
