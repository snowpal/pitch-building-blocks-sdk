package response

import (
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
)

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
