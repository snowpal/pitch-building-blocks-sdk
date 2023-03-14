package request

import "development/go/recipes/lib/golang/structs/common"

type Scales struct {
	IncludeCounts *bool `json:"includeCounts"`
	ExcludeEmpty  *bool `json:"excludeEmpty"`
}

type Scale struct {
	ID           *string   `json:"id"`
	Name         *string   `json:"scaleName"`
	Type         *string   `json:"scaleType"`
	ScaleValues  *[]string `json:"scaleValues"`
	LastModified *string   `json:"lastModified"`

	Blocks *[]common.SlimBlock `json:"blocks"`
	Pods   *[]common.SlimPod   `json:"pods"`
}

type UpdateScaleValueReqBody struct {
	ScaleValue string `json:"scaleValue"`
}

type AddScaleIdParam struct {
	ScaleId string
	KeyId   string
	BlockId *string
	PodId   *string
}
