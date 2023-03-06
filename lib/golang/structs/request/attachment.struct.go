package request

type AttachmentsAdd struct {
	Attachments []AttachmentAdd `json:"files"`
}

type AttachmentAdd struct {
	FileName string `json:"fileName"`
	FileUrl  string `json:"fileUrl"`
}

type AttachmentUpdate struct {
	ID       string `json:"id"`
	FileName string `json:"fileName"`
}
