package main

import (
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TestWhatsappMessage_ParseFromNumber(t *testing.T) {
	type fields struct {
		Model           gorm.Model
		WhatsappId      string
		RemoteJid       string
		SenderJid       string
		FromMe          bool
		Timestamp       uint64
		PushName        string
		Status          int
		QuotedMessageID string
		Text            string
	}
	tests := []struct {
		name     string
		fields   fields
		wantFrom string
	}{
		{
			name:     "Test standard wa",
			fields:   fields{SenderJid: "6285219132738@s.whatsapp.net"},
			wantFrom: "+6285219132738",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			waMsg := &WhatsappMessage{
				Model:           tt.fields.Model,
				WhatsappId:      tt.fields.WhatsappId,
				RemoteJid:       tt.fields.RemoteJid,
				SenderJid:       tt.fields.SenderJid,
				FromMe:          tt.fields.FromMe,
				Timestamp:       tt.fields.Timestamp,
				PushName:        tt.fields.PushName,
				Status:          tt.fields.Status,
				QuotedMessageID: tt.fields.QuotedMessageID,
				Text:            tt.fields.Text,
			}
			if gotFrom := waMsg.ParseFromNumber(); gotFrom != tt.wantFrom {
				t.Errorf("WhatsappMessage.ParseFromNumber() = %v, want %v", gotFrom, tt.wantFrom)
			}
		})
	}
}
