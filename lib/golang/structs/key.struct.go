package structs

import "development/go/recipes/lib/golang/structs/common"

type SlimKey struct {
	ID   string `json:"id"`
	Name string `json:"keyName"`
	Type string `json:"keyType"`
}

type KeyCounts struct {
	BlocksCount     *int `json:"blocksCount"`
	PodsCount       *int `json:"podsCount"`
	TasksCount      *int `json:"tasksCount"`
	ChecklistsCount *int `json:"checklistsCount"`
	NotesCount      *int `json:"notesCount"`
}

type Key struct {
	slimKey     SlimKey
	Description string `json:"keyDescription"`
	commonAttrs common.ResourceAttributes
	counts      KeyCounts
}
