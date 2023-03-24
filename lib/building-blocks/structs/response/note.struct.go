package response

import "development/go/recipes/lib/building-blocks/structs/common"

type Notes struct {
	Notes []Note `json:"notes"`
}

type Note struct {
	ID           string                  `json:"id"`
	NoteText     string                  `json:"noteText"`
	Creator      common.ResourceCreator  `json:"creator"`
	Modifier     common.ResourceModifier `json:"modifier"`
	LastModified string                  `json:"lastModified"`
}
