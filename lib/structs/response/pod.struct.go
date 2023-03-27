package response

import (
	common2 "github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
)

type Pods struct {
	Pods []Pod `json:"pods"`
}

type Pod struct {
	ID                string  `json:"id"`
	Name              string  `json:"podName"`
	Description       string  `json:"podDescription"`
	SimpleDescription string  `json:"simpleDescription"`
	Color             string  `json:"color"`
	Tags              string  `json:"tags"`
	ScaleValue        *string `json:"scaleValue"`

	Attributes  []common2.DisplayAttribute `json:"attributes"`
	PodType     *PodType                   `json:"podType"`
	Scale       *Scale                     `json:"scale"`
	TaggedUsers []TaggedUser               `json:"taggedUsers"`
	Key         *common2.SlimKey           `json:"key"`
	Block       *common2.SlimBlock         `json:"block"`

	// Boolean Attributes
	Completed    *bool `json:"completed"`
	Archived     *bool `json:"archived"`
	KanbanMode   *bool `json:"kanbanMode"`
	CanUnlink    *bool `json:"canUnlink"`
	CanUnpublish bool  `json:"canUnpublish"`
	PublicKey    *bool `json:"publicKey"`

	// Project Key Attribute
	ProjectKanbanMode *bool   `json:"projectKanbanMode"`
	ProjectListName   *string `json:"projectListName"`

	// Time Attributes
	DueDate string `json:"podDueDate"`

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
	BlocksCount      *int `json:"podsCount"`
	TasksCount       *int `json:"tasksCount"`
	ChecklistsCount  *int `json:"checklistsCount"`
	AttachmentsCount *int `json:"attachmentsCount"`
	CommentsCount    *int `json:"commentsCount"`
	NotesCount       *int `json:"notesCount"`

	Creator      common2.ResourceCreator  `json:"creator"`
	Modifier     common2.ResourceModifier `json:"modifier"`
	LastModified string                   `json:"lastModified"`
}

type UpdatePodScaleValue struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	ScaleValue   string `json:"scaleValue"`
	NumericScale int    `json:"numericScale"`

	Key   common2.SlimKey    `json:"key"`
	Block *common2.SlimBlock `json:"block"`
}
