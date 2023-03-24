package response

import "development/go/recipes/lib/building-blocks/structs/common"

type BlockTypes struct {
	BlockTypes []BlockType `json:"blockTypes"`
}

type BlockType struct {
	ID     string              `json:"id"`
	Name   string              `json:"blockTypeName"`
	Blocks *[]common.SlimBlock `json:"blocks"`

	TeacherReadOnly *bool `json:"teacherReadOnly"`

	Modifier     *common.ResourceModifier `json:"modifier"`
	LastModified string                   `json:"lastModified"`
}
