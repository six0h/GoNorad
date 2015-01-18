package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "encoding/json"
    "github.com/six0h/neptune"
)

type Error struct {
    code    int
    message string
}

type Config struct {
    username        string
    password        string
}

type FleetType struct {
    ships           int
    owner_id        int
    destination_id  int
    name            string
}

type PlayerType struct {
    id              int
    name            string
    economy         int
    industry        int
    science         int
    stars           int
}

type StarType struct {
    id              int
    name            string
    economy         int
    industry        int
    science         int
}

var config = make(map[string]string)

func main() {

    file, e := ioutil.ReadFile("app.config")
    if e != nil {
        fmt.Printf("File Error: %v\n", e)
        os.Exit(1)
    }

    json.Unmarshal(file, config)

    neptune.Login(config["username"], config["password"])
    var data = neptune.GetData()
    fmt.Println(data)

}
