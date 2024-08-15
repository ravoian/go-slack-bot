package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

func main() {
	// Configure Slack data
	url := "https://slack.com/api/chat.postMessage"
	contentType := "application/json"
	channelId := "<your_channel_here>"
	message := "go"
	token := "<your_token_here>"
	bearer := fmt.Sprintf("Bearer %s", token)
	datas := fmt.Sprintf(`{"channel": "%s", "text": "%s"}`, channelId, message)
	data := []byte(datas)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	// Configure headers
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Authorization", bearer)

	// Error handling
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Display result
	json_msg := string(body)
	buf := new(bytes.Buffer)
	json.Indent(buf, []byte(json_msg), "", "  ")
	fmt.Println(buf)
}