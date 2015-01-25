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
	var np2 Neptune
	config := GetConfig()
	cookieJar, _ := cookiejar.New(nil)

	_, e := np2.Login(config["username"], config["password"], cookieJar)
	errorExit(e)

	data, err := np2.GetData(config["gameNumber"], cookieJar)
	errorExit(err)

	fmt.Println(data)
}

func errorExit(e error) {
	if e != nil {
		fmt.Println("Error:", e)
		os.Exit(1)
	}
}
