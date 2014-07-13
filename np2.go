package neptune

import (
    "strconv"
    "net/http"
    "net/url"
    "bytes"
)

const (
    BASE_URL = "http://triton.ironhelmet.com"
    LOGIN_RESOURCE = "/a/login"
    ORDER_RESOURCE = "/grequest/order"
)

type ReportType struct {
    player_id       int
    fleets          map[string]FleetType
    players         map[string]PlayerType
    stars           map[string]StarType
}

type NeptuneResponse struct {
    event           string
    order           string
    error           string
}

func Login(username string, password string) {

    data := url.Values{}
    data.Set("username", username)
    data.Add("password", password)

    u, _ := url.ParseRequestURI(BASE_URL)
    u.Path = LOGIN_RESOURCE
    urlStr := fmt.Sprintf("%v", u)

    client := &http.Client{}
    r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
    r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

    client.Do(r)

}

func GetData() string {

    orderType := "full_universe_report"
    requestType := "order"
    version := "7"
    gameNumber := "5664221551394816"

    data := url.Values{}
    data.Set("order", orderType)
    data.Add("type", requestType)
    data.Add("version", version)
    data.Add("game_number", gameNumber)

    u, _ := url.ParseRequestURI(BASE_URL)
    u.Path = ORDER_RESOURCE
    urlStr := fmt.Sprintf("%v", u)

    client := &http.Client{}
    r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
    r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

    res, err := client.Do(r)
    if err != nil {
        fmt.Println(err) 
    }

    resp, err := ioutil.ReadAll(res.Body)

    return string(resp)
}
