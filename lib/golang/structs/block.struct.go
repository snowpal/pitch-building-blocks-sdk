package structs

import "development/go/recipes/lib/golang/structs/common"

type SlimBlock struct {
	ID   string `json:"id"`
	Name string `json:"blockName"`
	key  SlimKey
}

type Block struct {
	slimBlock   SlimBlock
	Description string `json:"blockDescription"`
	commonAttrs common.ResourceAttributes
	CanUnlink   bool `json:"canUnlink"`
	shareable   common.ShareableResource
	counts      BlockCounts
}

type BlockCounts struct {
	PodsCount *int `json:"podsCount"`
	counts    common.BlockAndPodCounts
}
