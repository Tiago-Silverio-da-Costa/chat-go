package handlers

type Message struct {
	ClientName string `json:"client_name"`
	Text       string `json:"text"`
}