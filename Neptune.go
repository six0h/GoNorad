package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const DOMAIN string = "triton.ironhelmet.com"
const BASE_URL string = "http://triton.ironhelmet.com"
const LOGIN_RESOURCE string = "/arequest/login"
const ORDER_RESOURCE string = "/grequest/order"

type Neptune struct {
}

func (Neptune) Login(username string, password string, cookieJar *cookiejar.Jar) (body string, err error) {
	data := url.Values{
		"alias":    {username},
		"password": {password},
		"type":     {"login"},
	}

	body, err = doRequest(LOGIN_RESOURCE, data, cookieJar)

	return
}

func (Neptune) GetData(gameNumber string, cookieJar *cookiejar.Jar) (body string, err error) {
	data := url.Values{
		"order":       {"full_universe_report"},
		"type":        {"order"},
		"version":     {"7"},
		"game_number": {gameNumber},
	}

	body, err = doRequest(ORDER_RESOURCE, data, cookieJar)

	return
}

func doRequest(resource string, data url.Values, cookieJar *cookiejar.Jar) (body string, e error) {
	u, _ := url.ParseRequestURI(BASE_URL)
	u.Path = resource
	urlStr := fmt.Sprintf("%v", u)

	var client = &http.Client{
		Jar: cookieJar,
	}

	r, _ := client.PostForm(urlStr, data)
	response, e := ioutil.ReadAll(r.Body)

	for cookie := range cookieJar.Cookies(u) {
		fmt.Println(cookie)
	}

	body = string(response)

	return
}
