package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Push_Receive(t *testing.T) {
	scenarios := []struct {
		receiveMessage   Message
		expectedResponse PageResponse
	}{
		{
			receiveMessage:   Message{From: "+6285219132737", To: "+6289602812374", Message: "halo halo"},
			expectedResponse: PageResponse{HTTPCode: 201, Data: Message{From: "+6285219132737", To: "+6289602812374", Message: "halo halo"}, ErrorMessage: ""},
		},
	}
	for _, scenario := range scenarios {

		t.Run("Run Push", func(t *testing.T) {
			push := PushReceive(scenario.receiveMessage)
			assert.Equal(t, scenario.expectedResponse, push)
		})

	}
}
