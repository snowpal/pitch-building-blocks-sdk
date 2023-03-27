package response

import (
	common2 "github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
)

type Checklists struct {
	Checklists []Checklist `json:"checklists"`
}

type Checklist struct {
	ID             string                   `json:"id"`
	Title          string                   `json:"title"`
	ChecklistItems []ChecklistItem          `json:"checklistItems"`
	Key            *common2.SlimKey         `json:"key"`
	Creator        common2.ResourceCreator  `json:"creator"`
	Modifier       common2.ResourceModifier `json:"modifier"`
}

type ChecklistItems struct {
	ChecklistItems []ChecklistItem `json:"checklistItems"`
}

type ChecklistItem struct {
	ID          string       `json:"id"`
	ItemTitle   string       `json:"title"`
	Completed   bool         `json:"completed"`
	TaggedUsers []TaggedUser `json:"taggedUsers"`

	ItemSequence *int                      `json:"checklistItemSeq"`
	Creator      *common2.ResourceCreator  `json:"creator"`
	Modifier     *common2.ResourceModifier `json:"modifier"`
}
