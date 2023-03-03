package structs

import "development/go/recipes/lib/golang/structs/common"

type TeacherPod struct {
	slimPod      SlimBlockPod
	Description  string `json:"podDescription"`
	commonAttrs  common.ResourceAttributes
	Published    bool `json:"completed"`
	CanUnlink    bool `json:"canUnlink"`
	CanUnpublish bool `json:"canUnpublish"`
	counts       BlockPodCounts
}
