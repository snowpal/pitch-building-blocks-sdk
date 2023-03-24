package response

import "development/go/recipes/lib/building-blocks/structs/common"

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

	Key    *common.SlimKey     `json:"key"`
	Block  *common.SlimBlock   `json:"block"`
	Blocks *[]common.SlimBlock `json:"blocks"`

	Modifier common.ResourceModifier `json:"modifier"`
}

type Relations struct {
	Relationships Relationships `json:"relationships"`
}

type Relationships struct {
	Keys   *common.SlimKey   `json:"keys"`
	Blocks *common.SlimBlock `json:"blocks"`
	Pods   *common.SlimPod   `json:"pods"`
}
