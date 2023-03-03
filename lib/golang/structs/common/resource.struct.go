package common

type DisplayAttribute struct {
	Name string `json:"attributeName"`
	Show bool   `json:"show"`
}

type ResourceCreator struct {
	CreatedOn string `json:"createdOn"`
	CreatedBy string `json:"createdBy"`
}

type ResourceModifier struct {
	ModifiedOn string `json:"modifiedOn"`
	ModifiedBy string `json:"modifiedBy"`
}

type ResourceAttributes struct {
	Creator           ResourceCreator    `json:"creator"`
	Modifier          ResourceModifier   `json:"modifier"`
	Attributes        []DisplayAttribute `json:"attributes"`
	SimpleDescription string             `json:"simpleDescription"`
	Color             string             `json:"color"`
	Tags              string             `json:"tags"`
	Archived          bool               `json:"archived"`
	KanbanMode        bool               `json:"kanbanMode"`
}

type ShareableResource struct {
	IsShared    bool   `json:"isShared"`
	Acl         string `json:"acl"`
	CanLeave    bool   `json:"canLeave"`
	AllowDelete bool   `json:"allowDelete"`
	CanDelete   bool   `json:"canDelete"`
}

type BlockAndPodCounts struct {
	KeysCount        *int `json:"keysCount"`
	PodsCount        *int `json:"podsCount"`
	TasksCount       *int `json:"tasksCount"`
	ChecklistsCount  *int `json:"checklistsCount"`
	AttachmentsCount *int `json:"attachmentsCount"`
	CommentsCount    *int `json:"commentsCount"`
	NotesCount       *int `json:"notesCount"`
}
