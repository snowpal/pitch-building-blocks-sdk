package response

import (
	"development/go/recipes/lib/golang/structs/common"
)

type Blocks struct {
	Blocks []Block `json:"blocks"`
}

type Block struct {
	ID                string `json:"id"`
	Name              string `json:"blockName"`
	BlockId           string `json:"blockId"`
	Description       string `json:"blockDescription"`
	SimpleDescription string `json:"simpleDescription"`
	Color             string `json:"color"`
	Tags              string `json:"tags"`

	Attributes  []common.DisplayAttribute `json:"attributes"`
	BlockType   *BlockType                `json:"blockType"`
	Scale       *Scale                    `json:"scale"`
	TaggedUsers []TaggedUser              `json:"taggedUsers"`
	Key         *common.SlimKey           `json:"key"`

	// Boolean Attributes
	Completed  *bool `json:"completed"`
	Archived   *bool `json:"archived"`
	KanbanMode *bool `json:"kanbanMode"`
	CanUnlink  *bool `json:"canUnlink"`
	PublicKey  *bool `json:"publicKey"`

	// Project Key Attribute
	ProjectKanbanMode *bool `json:"projectKanbanMode"`

	// Time Attributes
	DueDate   string `json:"blockDueDate"`
	StartTime string `json:"blockStartTime"`
	EndTime   string `json:"blockEndTime"`

	// Acl Attributes
	Acl            *string       `json:"acl"`
	IsShared       *bool         `json:"isShared"`
	CanLeave       *bool         `json:"canLeave"`
	AllowDelete    *bool         `json:"allowDelete"`
	CanDelete      *bool         `json:"canDelete"`
	SharedUsers    *[]SharedUser `json:"sharedUsers"`
	CurrentUserAcl SharedUser    `json:"currentUserAcl"`

	// Count Attributes
	KeysCount        *int `json:"keysCount"`
	PodsCount        *int `json:"podsCount"`
	TasksCount       *int `json:"tasksCount"`
	ChecklistsCount  *int `json:"checklistsCount"`
	AttachmentsCount *int `json:"attachmentsCount"`
	CommentsCount    *int `json:"commentsCount"`
	NotesCount       *int `json:"notesCount"`

	Creator      common.ResourceCreator  `json:"creator"`
	Modifier     common.ResourceModifier `json:"modifier"`
	LastModified string                  `json:"lastModified"`
}