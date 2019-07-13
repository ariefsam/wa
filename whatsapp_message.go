package main

import (
	"strings"

	"github.com/Rhymen/go-whatsapp"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type WhatsappMessage struct {
	gorm.Model
	WhatsappId      string `gorm:"type:varchar(255);unique_index"`
	RemoteJid       string
	SenderJid       string
	FromMe          bool
	Timestamp       uint64
	PushName        string
	Status          int
	QuotedMessageID string
	Text            string `gorm:"type:longtext;unique_index"`
}

func (waMsg *WhatsappMessage) Parse(wa whatsapp.TextMessage) {
	waMsg.WhatsappId = wa.Info.Id
	waMsg.PushName = wa.Info.PushName
	waMsg.Text = wa.Text
	waMsg.RemoteJid = wa.Info.RemoteJid
	waMsg.SenderJid = wa.Info.SenderJid
	if waMsg.SenderJid == "" {
		waMsg.SenderJid = wa.Info.Source.GetParticipant()
	}
	if waMsg.SenderJid == "" {
		waMsg.SenderJid = waMsg.RemoteJid
	}
	waMsg.FromMe = wa.Info.FromMe
	waMsg.Timestamp = wa.Info.Timestamp
	waMsg.Status = int(wa.Info.Status)
	waMsg.QuotedMessageID = wa.Info.QuotedMessageID
}

func (waMsg *WhatsappMessage) Save() error {
	var b error
	a := func(db *gorm.DB) {
		err := db.Create(&waMsg).Error
		b = err
	}
	DB(a)
	return b
}

func (waMsg *WhatsappMessage) FindByWhatsappID(id string) {

	a := func(db *gorm.DB) {
		db.First(&waMsg, "whatsapp_id=?", id)
	}
	DB(a)
	return
}

func (waMsg *WhatsappMessage) ParseFromNumber() (from string) {
	stringSlice := strings.Split(waMsg.SenderJid, "@")
	from = "+" + stringSlice[0]
	return
}

func (waMsg *WhatsappMessage) IsGroup() (ret bool) {
	stringSlice := strings.Split(waMsg.RemoteJid, "@")
	if len(stringSlice) > 1 {
		if stringSlice[1] == "g.us" {
			ret = true
		}
	}
	return
}
