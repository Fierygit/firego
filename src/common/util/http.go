package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func MakeSlackBotReq(text string) {
	url := "https://hooks.slack.com/services/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"

	payload := strings.NewReader(fmt.Sprintf("{\"text\": \"%s\"}", text))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("user-agent", "golang-client")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	ioutil.ReadAll(res.Body)
}
