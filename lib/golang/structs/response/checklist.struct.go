package response

import "development/go/recipes/lib/golang/structs/common"

type Checklists struct {
	Checklists []Checklist `json:"checklists"`
}

type Checklist struct {
	ID             string                  `json:"id"`
	Title          string                  `json:"checklistTitle"`
	Key            *common.SlimKey         `json:"key"`
	Creator        common.ResourceCreator  `json:"creator"`
	Modifier       common.ResourceModifier `json:"modifier"`
	ChecklistItems []ChecklistItem         `json:"checklistItems"`
}

type ChecklistItem struct {
	ID          string `json:"id"`
	ItemText    string `json:"checklistItemText"`
	Completed   bool   `json:"completed"`
	TaggedUsers []User `json:"taggedUsers"`
}
