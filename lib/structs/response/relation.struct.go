package response

import (
	common2 "github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
)

type SearchResources struct {
	Results []SearchResource `json:"results"`
}

type SearchResource struct {
	ID   string `json:"id"`
	Type string `json:"type"`

	// Relation attribute
	IsRelated bool `json:"isRelated"`

	KeyName   *string `json:"keyName"`
	KeyType   *string `json:"keyType"`
	BlockName *string `json:"blockName"`
	PodName   *string `json:"podName"`

	Key    *common2.SlimKey     `json:"key"`
	Block  *common2.SlimBlock   `json:"block"`
	Blocks *[]common2.SlimBlock `json:"blocks"`

	Modifier common2.ResourceModifier `json:"modifier"`
}

type Relations struct {
	Relationships Relationships `json:"relationships"`
}

type Relationships struct {
	Keys   []common2.SlimKey   `json:"keys"`
	Blocks []common2.SlimBlock `json:"blocks"`
	Pods   []common2.SlimPod   `json:"pods"`
}
