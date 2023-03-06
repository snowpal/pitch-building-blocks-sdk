package request

type ConversationAdd struct {
	MessageText string `json:"messageText"`
	Usernames   string `json:"userNames"`
}

type ConversationUpdate struct {
	ID          string `json:"id"`
	MessageText string `json:"messageText"`
}
