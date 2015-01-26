package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetConfig() map[string]string {

	binPath, err := GetCurrentPath()
	ErrorExit(err)

	data, err := ioutil.ReadFile(binPath + "/config.json")
	ErrorExit(err)

	var config = make(map[string]string)
	json.Unmarshal(data, &config)
	return config
}

func GetCurrentPath() (dir string, err error) {
	dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		ErrorExit(err)
	}

	return
}
