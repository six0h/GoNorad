package main

import (
	"fmt"
	"net/http/cookiejar"
	"net/url"
)

type DataRequest struct {
	result []string
}

func main() {
	config := GetConfig()

	cookieJar, _ := cookiejar.New(nil)

	_, e := Login(config["username"], config["password"], &cookieJar)
	if e != nil {
		fmt.Println("Error:")
		fmt.Println(e)
	}

	address, _ := url.Parse(BASE_URL)
	for cookie := range cookieJar.Cookies(address) {
		fmt.Println("Cookie")
		fmt.Println(cookie)
	}

	data, err := GetData(config["gameNumber"], &cookieJar)
	if err != nil {
		fmt.Println("Error Getting Data:")
		fmt.Println(err)
	}

	fmt.Println(data)
}
