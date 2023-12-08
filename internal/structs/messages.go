package structs

type Message struct {
	IsUserMessage bool   `json:"is_user_message"`
	Text          string `json:"text"`
}
