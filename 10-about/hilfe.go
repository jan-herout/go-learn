package main

/*
   Create a new Slack app in the workspace where you want to post messages.
   From the Features page, toggle Activate Incoming Webhooks on.
   Click Add New Webhook to Workspace.
   Pick a channel that the app will post to, then click Authorize.
   Use your Incoming Webhook URL to post a message to Slack.
*/

import (
	"bytes"
	"os"
	"fmt"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

type SlackRequestBody struct {
	Text string `json:"text"`
}


func main() {
	 os.Setenv("http_proxy", "")
	 os.Setenv("https_proxy", "")
	// START,OMIT
	hook := os.Getenv("SLACK_WEBHOOK")	
	msg := "Máchale, spadlo ti to"	
	err := SendSlackNotification(hook, msg)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	// END,OMIT
}

// SendSlackNotification will post to an 'Incoming Webook' url setup in Slack Apps. It accepts
// some text and the slack channel is saved within Slack.
func SendSlackNotification(webhookUrl string, msg string) error {

	slackBody, _ := json.Marshal(SlackRequestBody{Text: msg})
	fmt.Println("WEBHOOK=",webhookUrl)
	fmt.Println("MSG=",string(slackBody))
	req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	
	client := &http.Client{Timeout: 2 * time.Second}
	fmt.Println("Sending request...")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack")
	}
	return nil
}
