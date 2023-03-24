package response

import "development/go/recipes/lib/building-blocks/structs/common"

type PodTypes struct {
	PodTypes []PodType `json:"podTypes"`
}

type PodType struct {
	ID   string            `json:"id"`
	Name string            `json:"podTypeName"`
	Pods *[]common.SlimPod `json:"pods"`

	TeacherReadOnly *bool `json:"teacherReadOnly"`

	Modifier     common.ResourceModifier `json:"modifier"`
	LastModified string                  `json:"lastModified"`
}
