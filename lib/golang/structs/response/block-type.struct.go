package response

import "development/go/recipes/lib/golang/structs/common"

type BlockTypes struct {
	BlockTypes []BlockType `json:"blockTypes"`
}

type BlockType struct {
	ID              string                  `json:"id"`
	Name            string                  `json:"blockTypeName"`
	TeacherReadOnly *bool                   `json:"teacherReadOnly"`
	Modifier        common.ResourceModifier `json:"modifier"`
	LastModified    string                  `json:"lastModified"`
}
