package request

type ConversationUpdate struct {
	ID          *string `json:"id"`
	MessageText *string `json:"messageText"`
	Usernames   *string `json:"userNames"`
}
