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
