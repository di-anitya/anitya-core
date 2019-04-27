package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "os"
    "strings"
)

// IncomingURL - Get it from here https://slack.com/services/new/incoming-webhook
var IncomingURL string = "https://hooks.slack.com/services/XXX/XXX/XXX"

// Slack struct - payload parameter of json to post.
type Slack struct {
    Text      string `json:"text"`
    Username  string `json:"username"`
    IconEmoji string `json:"icon_emoji"`
    IconURL   string `json:"icon_url"`
    Channel   string `json:"channel"`
}

func main() {
    arg := strings.Join(os.Args[1:], "")
    params := Slack{
        Text:      fmt.Sprintf("%s", arg),
        Username:  "From golang to slack hello",
        IconEmoji: ":gopher:",
        IconURL:   "",
        Channel:   "",
    }
    jsonparams, _ := json.Marshal(params)
    resp, _ := http.PostForm(
        IncomingURL,
        url.Values{"payload": {string(jsonparams)}},
    )
    body, _ := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()

    println(string(body))
}