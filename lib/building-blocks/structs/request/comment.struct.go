package request

type CommentReqBody struct {
	CommentText   string  `json:"commentText"`
	TaggedUserIds *string `json:"taggedUserIds"`
}

type CommentIdParam struct {
	KeyId     string
	BlockId   *string
	PodId     *string
	CommentId *string
}
