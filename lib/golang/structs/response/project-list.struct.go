package response

import "development/go/recipes/lib/golang/structs/common"

type ProjectLists struct {
	ProjectLists []ProjectList `json:"projectLists"`
}

type ProjectList struct {
	ID       string `json:"id"`
	Name     string `json:"projectListName"`
	Sequence int    `json:"projectListSequence"`

	Key   common.SlimKey `json:"key"`
	Block common.SlimPod `json:"block"`
	Pods  []Pod          `json:"pods"`

	Creator      common.ResourceModifier `json:"creator"`
	Modifier     common.ResourceModifier `json:"modifier"`
	LastModified string                  `json:"lastModified"`
}
