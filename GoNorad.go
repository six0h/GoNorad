package main

import (
	"fmt"
	"github.com/six0h/neptune"
	"io/ioutil"
	"net/http/cookiejar"
	"net/url"
)

func main() {
	config := GetConfig()
	cookieJar, _ := cookiejar.New(nil)

	fmt.Println("Config:")
	fmt.Println(config)
	fmt.Println(config["username"])
	fmt.Println(config["password"])
	fmt.Println(config["gameNumber"])

	r, e := neptune.Login(config["username"], config["password"], &cookieJar)
	if e != nil {
		fmt.Println("Error:")
		fmt.Println(e)
	}

	loginResponse, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(loginResponse))

	address, _ := url.Parse(neptune.BASE_URL)
	for cookie := range cookieJar.Cookies(address) {
		fmt.Println("Cookie")
		fmt.Println(cookie)
	}

	data, err := neptune.GetData(config["gameNumber"], &cookieJar)
	if err != nil {
		fmt.Println("Error Getting Data:")
		fmt.Println(err)
	}

	fmt.Println(data)
}
