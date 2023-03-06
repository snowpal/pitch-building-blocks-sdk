package response

import "development/go/recipes/lib/golang/structs/common"

type Relations struct {
	Relations Relationships `json:"relationships"`
}

type Relationships struct {
	Keys   *common.SlimKey   `json:"keys"`
	Blocks *common.SlimBlock `json:"blocks"`
	Pods   *common.SlimPod   `json:"pods"`
}
