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

func Login(username string, password string, cookieJar **cookiejar.Jar) (res []byte, e error) {
	u, _ := url.ParseRequestURI(BASE_URL)
	u.Path = LOGIN_RESOURCE
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{
		Jar: *cookieJar,
	}

	data := url.Values{
		"alias":    {username},
		"password": {password},
		"type":     {"login"},
	}

	req, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode()))
	req.Header.Set("Origin", "http://triton.ironhelmet.com")
	req.Header.Set("Referrer", "http://triton.ironhelmet.com/")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.85 Safari/537.36")
	req.Header.Set("Host", "triton.ironhelmet.com")
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "en-GB,en-US;q=0.8,en;q=0.6")
	req.Header.Set("Connection", "keep-alive")

	r, e := client.Do(req)
	defer r.Body.Close()

	res, _ = ioutil.ReadAll(r.Body)

	return
}

func GetData(gameNumber string, cookieJar **cookiejar.Jar) (body string, err error) {
	orderType := "full_universe_report"
	requestType := "order"
	version := "7"

	data := url.Values{}
	data.Set("order", orderType)
	data.Add("type", requestType)
	data.Add("version", version)
	data.Add("game_number", gameNumber)

	u, _ := url.ParseRequestURI(BASE_URL)
	u.Path = ORDER_RESOURCE
	urlStr := fmt.Sprintf("%v", u)

	var client = &http.Client{
		Jar: *cookieJar,
	}

	r, _ := client.PostForm(urlStr, data)
	response, err := ioutil.ReadAll(r.Body)

	body = string(response)

	return
}
