package response

import "development/go/recipes/lib/golang/structs/common"

type Attachments struct {
	Attachments []Attachment `json:"files"`
}

type Attachment struct {
	ID           string                  `json:"id"`
	FileName     string                  `json:"fileName"`
	FileUrl      string                  `json:"fileUrl"`
	FileSize     string                  `json:"fileSize"`
	Creator      common.ResourceCreator  `json:"creator"`
	Modifier     common.ResourceModifier `json:"modifier"`
	LastModified string                  `json:"lastModified"`
}
