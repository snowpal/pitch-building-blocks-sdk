package response

type Conversations struct {
	Conversations []Conversation `json:"conversations"`
}

type Conversation struct {
	ID            string `json:"id"`
	Type          string `json:"type"`
	RecentMessage string `json:"recentMessage"`

	IsGroup         bool `json:"isGroup"`
	UserBlocked     bool `json:"blockedUser"`
	UserLeft        bool `json:"leftConversation"`
	UserCanLeave    bool `json:"canLeave"`
	Archived        bool `json:"archived"`
	AllMessagesRead bool `json:"allMessagesRead"`

	Members      []Member               `json:"members"`
	Messages     *[]ConversationMessage `json:"messages"`
	LastModified string                 `json:"lastModified"`
}

type Member struct {
	ID                string `json:"id"`
	Username          string `json:"userName"`
	FirstName         string `json:"firstName"`
	FullName          string `json:"fullName"`
	LastSeenMessageId int    `json:"LastSeenMessageId"`
}

type ConversationMessage struct {
	ID          string `json:"id"`
	Sequence    int    `json:"sequence"`
	AddedBy     string `json:"addedBy"`
	MessageText string `json:"messageText"`
	MessageTime string `json:"messageTime"`
	ReadReceipt string `json:"readReceipt"`
}
