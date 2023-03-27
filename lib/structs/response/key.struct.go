package response

import (
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
)

type Keys struct {
	Keys []Key `json:"keys"`
}

type Key struct {
	ID                string                    `json:"id"`
	Name              string                    `json:"keyName"`
	Type              string                    `json:"keyType"`
	Description       string                    `json:"keyDescription"`
	SimpleDescription string                    `json:"simpleDescription"`
	Color             string                    `json:"color"`
	Tags              string                    `json:"tags"`
	Attributes        []common.DisplayAttribute `json:"attributes"`

	// Boolean Attributes
	Archived          *bool `json:"archived"`
	KanbanMode        *bool `json:"kanbanMode"`
	ProjectKanbanMode *bool `json:"projectKanbanMode"`
	Public            *bool `json:"public"`

	// Count Attributes
	BlocksCount     *int `json:"blocksCount"`
	PodsCount       *int `json:"podsCount"`
	TasksCount      *int `json:"tasksCount"`
	ChecklistsCount *int `json:"checklistsCount"`
	NotesCount      *int `json:"notesCount"`

	Creator      common.ResourceCreator  `json:"creator"`
	Modifier     common.ResourceModifier `json:"modifier"`
	LastModified string                  `json:"lastModified"`
}
