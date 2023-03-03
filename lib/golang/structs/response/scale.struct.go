package response

import "development/go/recipes/lib/golang/structs/common"

type Scale struct {
	ID              string                  `json:"id"`
	Name            string                  `json:"scaleName"`
	Type            string                  `json:"scaleType"`
	ScaleValues     []string                `json:"scaleValues"`
	Modifier        common.ResourceModifier `json:"modifier"`
	LastModified    string                  `json:"lastModified"`
	TeacherReadOnly *bool                   `json:"teacherReadOnly"`
}
