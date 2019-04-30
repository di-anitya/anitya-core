package notification

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
var IncomingURL = "https://hooks.slack.com/services/XXX/XXX/XXX"

// slack struct - payload parameter of json to post.
type slack struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	IconURL   string `json:"icon_url"`
	Channel   string `json:"channel"`
}

// SendSlack is ..
func SendSlack() {
	arg := strings.Join(os.Args[1:], "")
	params := slack{
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
