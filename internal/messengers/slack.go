package messengers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SlackMessenger struct{}

func (messenger *SlackMessenger) SendMessage(webhookUrl, title, description, url string) error {
	blocks := []map[string]interface{}{
		{
			"type": "header",
			"text": map[string]interface{}{
				"type":  "plain_text",
				"text":  fmt.Sprintf("New Service Desk issue: %s", title),
				"emoji": true,
			},
		},
		{
			"type": "section",
			"text": map[string]interface{}{
				"type": "mrkdwn",
				"text": fmt.Sprintf("*Description:*\n%s", description),
			},
		},
		{
			"type": "actions",
			"elements": []map[string]interface{}{
				{
					"type": "button",
					"text": map[string]interface{}{
						"type":  "plain_text",
						"text":  "View issue",
						"emoji": true,
					},
					"url":   url,
					"style": "primary",
				},
			},
		},
	}

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	body, err := json.Marshal(map[string]interface{}{"blocks": blocks})
	if err != nil {
		return err
	}

	res, err := client.Post(webhookUrl, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to post to slack, response status code: %d", res.StatusCode)
	}

	return nil
}
