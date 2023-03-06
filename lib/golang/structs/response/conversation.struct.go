package response

type Conversations struct {
	Conversations []Conversation `json:"conversations"`
}

type Conversation struct {
	ID              string                 `json:"id"`
	Type            string                 `json:"type"`
	LastModified    string                 `json:"lastModified"`
	IsGroup         bool                   `json:"isGroup"`
	UserBlocked     bool                   `json:"blockedUser"`
	UserLeft        bool                   `json:"leftConversation"`
	UserCanLeave    bool                   `json:"canLeave"`
	Archived        bool                   `json:"archived"`
	Members         []Member               `json:"members"`
	AllMessagesRead bool                   `json:"allMessagesRead"`
	RecentMessage   string                 `json:"recentMessage"`
	Messages        *[]ConversationMessage `json:"messages"`
}

type Member struct {
	ID                string `json:"id"`
	Username          string `json:"userName"`
	FirstName         string `json:"firstName"`
	FullName          string `json:"fullName"`
	LastSeenMessageId string `json:"LastSeenMessageId"`
}

type ConversationMessage struct {
	ID          string `json:"id"`
	Sequence    string `json:"sequence"`
	AddedBy     string `json:"addedBy"`
	MessageText string `json:"messageText"`
	MessageTime string `json:"messageTime"`
	ReadReceipt string `json:"readReceipt"`
}
