package request

type Blocks struct {
	Filter           *string `json:"filter"`
	BatchIndex       *int    `json:"batchIndex"`
	WriteOrHigherAcl *bool   `json:"aclWriteOrHigher"`
}

type Block struct {
	ID                *string `json:"id"`
	Name              *string `json:"blockName"`
	BlockId           *string `json:"blockId"`
	Description       *string `json:"blockDescription"`
	SimpleDescription *string `json:"simpleDescription"`
	DueDate           *string `json:"blockDueDate"`
	StartTime         *string `json:"blockStartTime"`
	EndTime           *string `json:"blockEndTime"`
	Color             *string `json:"blockColor"`
	Tags              *string `json:"blockTags"`
	ScaleValue        *string `json:"scaleValue"`

	Completed  *bool `json:"blockCompleted"`
	KanbanMode *bool `json:"kanbanMode"`

	TaggedUserIds *[]string `json:"taggedUserIds"`
	BlockIds      *[]string `json:"blockIds"`
}

type BlockByTemplate struct {
	TemplateId   string `json:"templateId"`
	ExcludeTasks *bool  `json:"excludeTasks"`
}

type ArchivedBlocks struct {
	KeyId      string `json:"keyId"`
	BatchIndex *int   `json:"batchIndex"`
}

type BlockWithParam struct {
	KeyId       string `json:"keyId"`
	TargetKeyId string `json:"targetKeyId"`

	PodIds        *[]string `json:"podIds"`
	AllPods       *bool     `json:"allPods"`
	AllTasks      *bool     `json:"allTasks"`
	AllChecklists *bool     `json:"allChecklists"`
}
