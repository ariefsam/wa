package main

import (
	"github.com/Rhymen/go-whatsapp"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type WhatsappMessage struct {
	gorm.Model
	WhatsappId      string `gorm:"type:varchar(100);unique_index"`
	RemoteJid       string
	SenderJid       string
	FromMe          bool
	Timestamp       uint64
	PushName        string
	Status          int
	QuotedMessageID string
	Text            string
}

func (waMsg *WhatsappMessage) Parse(wa whatsapp.TextMessage) {
	waMsg.WhatsappId = wa.Info.Id
	waMsg.PushName = wa.Info.PushName
	waMsg.Text = wa.Text
	waMsg.RemoteJid = wa.Info.RemoteJid
	waMsg.SenderJid = wa.Info.SenderJid
	waMsg.FromMe = wa.Info.FromMe
	waMsg.Timestamp = wa.Info.Timestamp
	waMsg.Status = int(wa.Info.Status)
	waMsg.QuotedMessageID = wa.Info.QuotedMessageID
}

func (waMsg *WhatsappMessage) Save() (err error) {
	a := func(db *gorm.DB) {

		db.Create(&waMsg)
	}
	DB(a)
	return
}

func (waMsg *WhatsappMessage) FindByWhatsappID(id string) {

	a := func(db *gorm.DB) {
		db.First(&waMsg, "whatsapp_id=?", id)
	}
	DB(a)
	return
}
