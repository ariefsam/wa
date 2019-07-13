package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Rhymen/go-whatsapp"
	"github.com/teris-io/shortid"
)

type sceWA struct {
	name          string
	wa            whatsapp.TextMessage
	expectedMesg  WhatsappMessage
	expectedGroup bool
}

func Test_waHandler_HandleTextMessage(t *testing.T) {

	var tests []sceWA
	var id string
	sid, _ := shortid.New(1, shortid.DefaultABC, 12)

	id, _ = sid.Generate()
	pushname, _ := sid.Generate()
	rndstring, _ := sid.Generate()

	tests = append(tests, sceWA{
		name: "Testing Normal",
		wa: whatsapp.TextMessage{
			Info: whatsapp.MessageInfo{
				Id:        id,
				PushName:  pushname,
				RemoteJid: "6285219132737@s.whatsapp.net",
				SenderJid: "",
			},
			Text: rndstring,
		},
		expectedMesg: WhatsappMessage{
			WhatsappId: id,
			PushName:   pushname,
			Text:       rndstring,
			RemoteJid:  "6285219132737@s.whatsapp.net",
			SenderJid:  "6285219132737@s.whatsapp.net",
		},
	})
	id, _ = sid.Generate()
	pushname, _ = sid.Generate()
	rndstring, _ = sid.Generate()
	tests = append(tests, sceWA{
		name: "Testing Group",
		wa: whatsapp.TextMessage{
			Info: whatsapp.MessageInfo{
				Id:        id,
				PushName:  pushname,
				RemoteJid: "6285219132737-1562921107@g.us",
				SenderJid: "6285219132737@s.whatsapp.net",
			},
			Text: rndstring,
		},
		expectedMesg: WhatsappMessage{
			WhatsappId: id,
			PushName:   pushname,
			Text:       rndstring,
			RemoteJid:  "6285219132737-1562921107@g.us",
			SenderJid:  "6285219132737@s.whatsapp.net",
		},

		expectedGroup: true,
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var msg WhatsappMessage
			msg.Parse(tt.wa)
			assert.Equal(t, tt.expectedMesg, msg)
			err := msg.Save()
			assert.NoError(t, err)

			var msgFromDB WhatsappMessage
			msgFromDB.FindByWhatsappID(tt.wa.Info.Id)
			tt.expectedMesg.ID = msgFromDB.ID
			tt.expectedMesg.CreatedAt = msgFromDB.CreatedAt
			tt.expectedMesg.UpdatedAt = msgFromDB.UpdatedAt
			assert.Equal(t, tt.expectedMesg, msgFromDB)
			assert.Equal(t, msgFromDB.IsGroup(), tt.expectedGroup)
		})
	}
}
