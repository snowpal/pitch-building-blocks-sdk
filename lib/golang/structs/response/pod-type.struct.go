package response

import "development/go/recipes/lib/golang/structs/common"

type PodTypes struct {
	PodTypes []PodType `json:"podTypes"`
}

type PodType struct {
	ID              string                  `json:"id"`
	Name            string                  `json:"podTypeName"`
	Modifier        common.ResourceModifier `json:"modifier"`
	LastModified    string                  `json:"lastModified"`
	TeacherReadOnly *bool                   `json:"teacherReadOnly"`
}
