package response

import (
	common2 "github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
)

type PodTypes struct {
	PodTypes []PodType `json:"podTypes"`
}

type PodType struct {
	ID   string             `json:"id"`
	Name string             `json:"podTypeName"`
	Pods *[]common2.SlimPod `json:"pods"`

	TeacherReadOnly *bool `json:"teacherReadOnly"`

	Modifier     common2.ResourceModifier `json:"modifier"`
	LastModified string                   `json:"lastModified"`
}
