package main

type Message struct {
	From      string
	To        string
	Message   string
	TimeStamp int64
}

type PageResponse struct {
	HTTPCode     int
	Data         Message
	ErrorMessage string
}
