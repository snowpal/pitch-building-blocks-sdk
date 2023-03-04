package response

import (
	"development/go/recipes/lib/golang/structs/common"
)

type Pods struct {
	Pods []Pod `json:"pods"`
}

type Pod struct {
	ID                string                    `json:"id"`
	PodClassType      string                    `json:"podClassType"`
	Name              string                    `json:"podName"`
	Description       string                    `json:"podDescription"`
	Key               *common.SlimKey           `json:"key"`
	Block             *common.SlimBlock         `json:"block"`
	Creator           common.ResourceCreator    `json:"creator"`
	Modifier          common.ResourceModifier   `json:"modifier"`
	LastModified      string                    `json:"lastModified"`
	Attributes        []common.DisplayAttribute `json:"attributes"`
	SimpleDescription string                    `json:"simpleDescription"`
	Color             string                    `json:"color"`
	Tags              string                    `json:"tags"`
	PodType           *PodType                  `json:"podType"`
	Scale             *Scale                    `json:"scale"`
	TaggedUsers       []User                    `json:"taggedUsers"`
	Archived          *bool                     `json:"archived"`
	KanbanMode        *bool                     `json:"kanbanMode"`
	ProjectKanbanMode *bool                     `json:"projectKanbanMode"`
	ProjectListName   *string                   `json:"projectListName"`
	Completed         *bool                     `json:"completed"`
	CanUnlink         *bool                     `json:"canUnlink"`
	CanUnpublish      bool                      `json:"canUnpublish"`
	IsShared          *bool                     `json:"isShared"`
	Acl               string                    `json:"acl"`
	CanLeave          *bool                     `json:"canLeave"`
	AllowDelete       *bool                     `json:"allowDelete"`
	CanDelete         *bool                     `json:"canDelete"`
	PublicKey         *bool                     `json:"publicKey"`
	Keys              int                       `json:"keysCount"`
	Blocks            int                       `json:"podsCount"`
	Tasks             int                       `json:"tasksCount"`
	Checklists        int                       `json:"checklistsCount"`
	Attachments       int                       `json:"attachmentsCount"`
	Comments          int                       `json:"commentsCount"`
	Notes             int                       `json:"notesCount"`
}
