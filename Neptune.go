package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

const DOMAIN string = "triton.ironhelmet.com"
const BASE_URL string = "http://triton.ironhelmet.com"
const LOGIN_RESOURCE string = "/arequest/login"
const ORDER_RESOURCE string = "/grequest/order"

type NeptuneRequest interface {
	DoRequest(string, url.Values, cookiejar.Jar) (string, error)
	Login(string, string, cookiejar.Jar) (string, error)
	GetData(string, cookiejar.Jar) (string, error)
}

type Neptune struct{}

func (Neptune) Login(username string, password string, cookieJar *cookiejar.Jar) (jsonBody string, err error) {
	data := url.Values{
		"alias":    {username},
		"password": {password},
		"type":     {"login"},
	}

	jsonBody, err = DoRequest(LOGIN_RESOURCE, data, cookieJar)
	ErrorExit(err)

	return
}

func (Neptune) GetData(gameNumber string, cookieJar *cookiejar.Jar) (jsonBody string, err error) {
	data := url.Values{
		"order":       {"full_universe_report"},
		"type":        {"order"},
		"version":     {"7"},
		"game_number": {gameNumber},
	}

	jsonBody, err = DoRequest(ORDER_RESOURCE, data, cookieJar)

	return
}

func DoRequest(resource string, data url.Values, cookieJar *cookiejar.Jar) (body string, e error) {
	u, _ := url.ParseRequestURI(BASE_URL)
	u.Path = resource
	urlStr := fmt.Sprintf("%v", u)

	var client = &http.Client{
		Jar: cookieJar,
	}

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	ErrorExit(err)

	r, e := client.Do(req)
	response, e := ioutil.ReadAll(r.Body)

	body = string(response)

	return
}
