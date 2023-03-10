package request

type AttachmentsReqBody struct {
	Attachments []AttachmentReqBody `json:"files"`
}

type AttachmentReqBody struct {
	FileName string  `json:"fileName"`
	FileUrl  *string `json:"fileUrl"`
}

type AttachmentParam struct {
	AttachmentId *string
	KeyId        string
	BlockId      *string
	PodId        *string
}
