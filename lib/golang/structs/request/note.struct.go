package request

type NoteReqBody struct {
	NoteText string `json:"noteText"`
}

type NoteIdParam struct {
	KeyId   string
	BlockId *string
	PodId   *string
	NoteId  *string
}
