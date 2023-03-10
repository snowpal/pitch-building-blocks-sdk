package request

type Attachments struct {
	Attachments []Attachment `json:"files"`
}

type Attachment struct {
	ID       *string `json:"id"`
	FileName *string `json:"fileName"`
	FileUrl  *string `json:"fileUrl"`
}

type AttachmentParam struct {
	AttachmentId string
	KeyId        string
	BlockId      *string
	PodId        *string
}
