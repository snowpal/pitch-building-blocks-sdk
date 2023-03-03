package structs

import "development/go/recipes/lib/golang/structs/common"

type SlimPod struct {
	ID   string `json:"id"`
	Name string `json:"podName"`
	key  SlimKey
}

type Pod struct {
	slimPod     SlimPod
	Description string `json:"podDescription"`
	commonAttrs common.ResourceAttributes
	Completed   bool `json:"completed"`
	CanUnlink   bool `json:"canUnlink"`
	shareable   common.ShareableResource
	counts      common.BlockAndPodCounts
}
