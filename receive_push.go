package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PushReceive(data Message) (response PageResponse) {
	requestBody, _ := json.Marshal(data)
	resp, err := http.Post(config.ReceivePushURL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &response)
	return
}
