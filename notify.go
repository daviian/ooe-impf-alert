package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type webhookBody struct {
	Date string `json:"value1"`
}

func sendPushNotification(date time.Time) error {
	body := &webhookBody{date.Format(time.RFC822)}

	buf := bytes.NewBuffer(make([]byte, 0))
	err := json.NewEncoder(buf).Encode(body)
	if err != nil {
		return fmt.Errorf("encode json failed: %w", err)
	}

	url := fmt.Sprintf("https://maker.ifttt.com/trigger/%s/with/key/%s", iftttEventName, iftttKey)
	_, err = http.Post(url, "application/json", buf)
	if err != nil {
		return fmt.Errorf("http post request failed: %w", err)
	}

	return nil
}
