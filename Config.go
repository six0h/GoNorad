package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func GetConfig() map[string]string {
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err)
	}
	var config = make(map[string]string)
	json.Unmarshal(data, &config)
	return config
}
