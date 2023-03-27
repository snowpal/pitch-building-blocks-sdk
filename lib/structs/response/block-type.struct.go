package response

import (
	common2 "github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
)

type BlockTypes struct {
	BlockTypes []BlockType `json:"blockTypes"`
}

type BlockType struct {
	ID     string               `json:"id"`
	Name   string               `json:"blockTypeName"`
	Blocks *[]common2.SlimBlock `json:"blocks"`

	TeacherReadOnly *bool `json:"teacherReadOnly"`

	Modifier     *common2.ResourceModifier `json:"modifier"`
	LastModified string                    `json:"lastModified"`
}
