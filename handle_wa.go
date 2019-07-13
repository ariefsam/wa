package main

import (
	"log"
	"time"

	"github.com/Rhymen/go-whatsapp"
)

type waHandler struct {
	c *whatsapp.Conn
}

//HandleError needs to be implemented to be a valid WhatsApp handler
func (h *waHandler) HandleError(err error) {

	if e, ok := err.(*whatsapp.ErrConnectionFailed); ok {
		log.Printf("Connection failed, underlying error: %v", e.Err)
		log.Println("Waiting 30sec...")
		<-time.After(30 * time.Second)
		log.Println("Reconnecting...")
		err := h.c.Restore()
		if err != nil {
			log.Fatalf("Restore failed: %v", err)
		}
	} else {
		log.Printf("error occoured: %v\n", err)
	}
}

//Optional to be implemented. Implement HandleXXXMessage for the types you need.
func (*waHandler) HandleTextMessage(message whatsapp.TextMessage) {
	//fmt.Printf("%v %v %v %v\n\t%v\n", message.Info.Timestamp, message.Info.Id, message.Info.RemoteJid, message.Info.QuotedMessageID, message.Text)
	var msg WhatsappMessage
	msg.Parse(message)
	err := msg.Save()
	if err == nil {
		if !msg.FromMe && !msg.IsGroup() {
			var receiveMessage Message
			receiveMessage.From = msg.ParseFromNumber()
			receiveMessage.Message = msg.Text
			receiveMessage.TimeStamp = msg.Timestamp
			PushReceive(receiveMessage)
		}
	}
}
