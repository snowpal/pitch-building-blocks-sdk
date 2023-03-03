package response

import (
	"development/go/recipes/lib/golang/structs/common"
)

type Keys struct {
	Keys []Key `json:"keys"`
}

type Key struct {
	ID                string                    `json:"id"`
	Name              string                    `json:"keyName"`
	Type              string                    `json:"keyType"`
	Description       string                    `json:"keyDescription"`
	Creator           common.ResourceCreator    `json:"creator"`
	Modifier          common.ResourceModifier   `json:"modifier"`
	LastModified      string                    `json:"lastModified"`
	Attributes        []common.DisplayAttribute `json:"attributes"`
	SimpleDescription string                    `json:"simpleDescription"`
	Color             string                    `json:"color"`
	Tags              string                    `json:"tags"`
	Archived          *bool                     `json:"archived"`
	KanbanMode        *bool                     `json:"kanbanMode"`
	ProjectKanbanMode *bool                     `json:"projectKanbanMode"`
	Public            *bool                     `json:"public"`
	Blocks            int                       `json:"blocksCount"`
	Pods              int                       `json:"podsCount"`
	Tasks             int                       `json:"tasksCount"`
	Checklists        int                       `json:"checklistsCount"`
	Notes             int                       `json:"notesCount"`
}
