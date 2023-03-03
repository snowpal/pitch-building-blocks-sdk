package structs

import "development/go/recipes/lib/golang/structs/common"

type ProjectPod struct {
	slimPod           SlimBlockPod
	Description       string `json:"podDescription"`
	commonAttrs       common.ResourceAttributes
	Completed         bool   `json:"completed"`
	CanUnlink         bool   `json:"canUnlink"`
	ProjectKanbanMode bool   `json:"projectKanbanMode"`
	ProjectListName   string `json:"projectListName"`
	counts            BlockPodCounts
}
