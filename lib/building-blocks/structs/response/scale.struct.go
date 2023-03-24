package response

import "development/go/recipes/lib/building-blocks/structs/common"

type Scales struct {
	Scales []Scale `json:"scales"`
}

type Scale struct {
	ID              string                  `json:"id"`
	Name            string                  `json:"scaleName"`
	Type            *string                 `json:"scaleType"`
	ScaleValues     []string                `json:"scaleValues"`
	TeacherReadOnly *bool                   `json:"teacherReadOnly"`
	Modifier        common.ResourceModifier `json:"modifier"`
	LastModified    string                  `json:"lastModified"`
}
