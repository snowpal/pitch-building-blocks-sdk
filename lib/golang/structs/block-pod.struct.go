package structs

import "development/go/recipes/lib/golang/structs/common"

type SlimBlockPod struct {
	ID    string `json:"id"`
	Name  string `json:"podName"`
	key   SlimKey
	block SlimBlock
}

type BlockPod struct {
	slimPod     SlimBlockPod
	Description string `json:"podDescription"`
	commonAttrs common.ResourceAttributes
	Completed   bool `json:"completed"`
	CanUnlink   bool `json:"canUnlink"`
	shareable   common.ShareableResource
	counts      BlockPodCounts
}

type BlockPodCounts struct {
	BlocksCount *int `json:"podsCount"`
	counts      common.BlockAndPodCounts
}
