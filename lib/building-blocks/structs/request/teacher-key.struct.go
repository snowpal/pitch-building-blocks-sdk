package request

import "development/go/recipes/lib/building-blocks/structs/common"

type PublishGradesReqBody struct {
	StudentUserIds string `json:"studentUserIds"`
}

type ClassroomIdParam struct {
	StudentId   string
	ResourceIds common.ResourceIdParam
}
