package response

import "development/go/recipes/lib/golang/structs/common"

type Checklists struct {
	Checklists []Checklist `json:"checklists"`
}

type Checklist struct {
	ID             string                  `json:"id"`
	Title          string                  `json:"checklistTitle"`
	ChecklistItems []ChecklistItem         `json:"checklistItems"`
	Key            *common.SlimKey         `json:"key"`
	Creator        common.ResourceCreator  `json:"creator"`
	Modifier       common.ResourceModifier `json:"modifier"`
}

type ChecklistItem struct {
	ID          string       `json:"id"`
	ItemText    string       `json:"checklistItemText"`
	Completed   bool         `json:"completed"`
	TaggedUsers []TaggedUser `json:"taggedUsers"`
}
