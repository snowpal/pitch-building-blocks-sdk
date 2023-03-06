package request

type CommentAdd struct {
	CommentText   string    `json:"commentText"`
	TaggedUserIds *[]string `json:"taggedUserIds"`
}

type CommentUpdate struct {
	ID            string    `json:"id"`
	CommentText   string    `json:"commentText"`
	TaggedUserIds *[]string `json:"taggedUserIds"`
}
