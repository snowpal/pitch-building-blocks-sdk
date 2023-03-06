package request

type Attachments struct {
	Attachments []Attachment `json:"files"`
}

type Attachment struct {
	ID       *string `json:"id"`
	FileName *string `json:"fileName"`
	FileUrl  *string `json:"fileUrl"`
}
