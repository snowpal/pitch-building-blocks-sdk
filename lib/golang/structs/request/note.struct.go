package request

type NoteAdd struct {
	NoteText string `json:"noteText"`
}

type NoteUpdate struct {
	ID       string `json:"id"`
	NoteText string `json:"noteText"`
}
