package response

type Comment struct {
	CommentText   string    `json:"commentText"`
	TaggedUserIds *[]string `json:"taggedUserIds"`
}
