package request

type Pods struct {
	KeyId      string `json:"keyId"`
	BatchIndex *int   `json:"batchIndex"`
}

type Pod struct {
	ID                *string `json:"id"`
	Name              *string `json:"podName"`
	Description       *string `json:"podDescription"`
	SimpleDescription *string `json:"simpleDescription"`
	DueDate           *string `json:"podDueDate"`
	Color             *string `json:"podColor"`
	Tags              *string `json:"podTags"`
	ScaleValue        *string `json:"scaleValue"`

	Completed  *bool `json:"podCompleted"`
	KanbanMode *bool `json:"kanbanMode"`

	TaggedUserIds *[]string `json:"taggedUserIds"`
	PodIds        *[]string `json:"podIds"`
}

type PodByTemplate struct {
	KeyId        string `json:"keyId"`
	TemplateId   string `json:"templateId"`
	ExcludeTasks *bool  `json:"excludeTasks"`
}

type ArchivedPods struct {
	KeyId      string  `json:"keyId"`
	BlockId    *string `json:"blockId"`
	BatchIndex *int    `json:"batchIndex"`
}

type PodWithParam struct {
	KeyId         string  `json:"keyId"`
	BlockId       *string `json:"blockId"`
	TargetKeyId   string  `json:"targetKeyId"`
	TargetBlockId *string `json:"targetBlockId"`

	AllTasks      *bool `json:"allTasks"`
	AllChecklists *bool `json:"allChecklists"`
}
