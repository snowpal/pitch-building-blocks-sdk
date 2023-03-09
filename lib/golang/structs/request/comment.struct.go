package request

type Comment struct {
	ID            *string   `json:"id"`
	CommentText   *string   `json:"commentText"`
	TaggedUserIds *[]string `json:"taggedUserIds"`
}
