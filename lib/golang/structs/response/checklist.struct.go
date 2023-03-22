package response

import "development/go/recipes/lib/golang/structs/common"

type Checklists struct {
	Checklists []Checklist `json:"checklists"`
}

type Checklist struct {
	ID             string                  `json:"id"`
	Title          string                  `json:"title"`
	ChecklistItems []ChecklistItem         `json:"checklistItems"`
	Key            *common.SlimKey         `json:"key"`
	Creator        common.ResourceCreator  `json:"creator"`
	Modifier       common.ResourceModifier `json:"modifier"`
}

type ChecklistItems struct {
	ChecklistItems []ChecklistItem `json:"checklistItems"`
}

type ChecklistItem struct {
	ID          string       `json:"id"`
	ItemTitle   string       `json:"title"`
	Completed   bool         `json:"completed"`
	TaggedUsers []TaggedUser `json:"taggedUsers"`

	ItemSequence *int                     `json:"checklistItemSeq"`
	Creator      *common.ResourceCreator  `json:"creator"`
	Modifier     *common.ResourceModifier `json:"modifier"`
}
