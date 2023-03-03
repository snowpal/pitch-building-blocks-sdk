package response

import (
	"development/go/recipes/lib/golang/structs/common"
)

type Blocks struct {
	Blocks []Block `json:"blocks"`
}

type Block struct {
	ID                string                    `json:"id"`
	BlockClassType    string                    `json:"blockClassType"`
	BlockId           string                    `json:"blockId"`
	Name              string                    `json:"blockName"`
	Description       string                    `json:"blockDescription"`
	Key               *common.SlimKey           `json:"key"`
	Creator           common.ResourceCreator    `json:"creator"`
	Modifier          common.ResourceModifier   `json:"modifier"`
	LastModified      string                    `json:"lastModified"`
	Attributes        []common.DisplayAttribute `json:"attributes"`
	SimpleDescription string                    `json:"simpleDescription"`
	DueDate           string                    `json:"blockDueDate"`
	StartTime         string                    `json:"blockStartTime"`
	EndTime           string                    `json:"blockEndTime"`
	Completed         *bool                     `json:"completed"`
	Color             string                    `json:"color"`
	Tags              string                    `json:"tags"`
	BlockType         *BlockType                `json:"blockType"`
	Scale             *Scale                    `json:"scale"`
	TaggedUsers       []User                    `json:"taggedUsers"`
	Archived          *bool                     `json:"archived"`
	KanbanMode        *bool                     `json:"kanbanMode"`
	ProjectKanbanMode *bool                     `json:"projectKanbanMode"`
	CanUnlink         *bool                     `json:"canUnlink"`
	IsShared          *bool                     `json:"isShared"`
	Acl               string                    `json:"acl"`
	CanLeave          *bool                     `json:"canLeave"`
	AllowDelete       *bool                     `json:"allowDelete"`
	CanDelete         *bool                     `json:"canDelete"`
	PublicKey         *bool                     `json:"publicKey"`
	Keys              int                       `json:"keysCount"`
	Pods              int                       `json:"podsCount"`
	Tasks             int                       `json:"tasksCount"`
	Checklists        int                       `json:"checklistsCount"`
	Attachments       int                       `json:"attachmentsCount"`
	Comments          int                       `json:"commentsCount"`
	Notes             int                       `json:"notesCount"`
}
