package response

import (
	common2 "github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
)

type ProjectLists struct {
	ProjectLists []ProjectList `json:"projectLists"`
}

type ProjectList struct {
	ID       string `json:"id"`
	Name     string `json:"projectListName"`
	Sequence int    `json:"projectListSequence"`

	Key   common2.SlimKey   `json:"key"`
	Block common2.SlimBlock `json:"block"`
	Pods  []Pod             `json:"pods"`

	Creator      common2.ResourceModifier `json:"creator"`
	Modifier     common2.ResourceModifier `json:"modifier"`
	LastModified string                   `json:"lastModified"`
}
