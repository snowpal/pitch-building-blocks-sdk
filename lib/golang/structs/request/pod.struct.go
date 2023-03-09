package request

type Pod struct {
	ID                *string `json:"id"`
	Name              *string `json:"podName"`
	Description       *string `json:"podDescription"`
	SimpleDescription *string `json:"simpleDescription"`
	DueDate           *string `json:"podDueDate"`
	Color             *string `json:"podColor"`
	Tags              *string `json:"podTags"`
	ScaleValue        *string `json:"scaleValue"`

	AllowArchival *bool `json:"allowArchival"`
	Completed     *bool `json:"podCompleted"`
	KanbanMode    *bool `json:"kanbanMode"`

	TaggedUserIds *[]string `json:"taggedUserIds"`
	PodIds        *[]string `json:"podIds"`
}

type PodByTemplate struct {
	KeyId        string `json:"keyId"`
	TemplateId   string `json:"templateId"`
	ExcludeTasks *bool  `json:"excludeTasks"`
}

type PodWithParam struct {
	TargetKeyId   string  `json:"targetKeyId"`
	TargetBlockId *string `json:"targetBlockId"`

	AllTasks      *bool `json:"allTasks"`
	AllChecklists *bool `json:"allChecklists"`
}

type SharePod struct {
	Acl            string    `json:"podAcl"`
	PodIds         *[]string `json:"podIds"`
	StudentUserIds *[]string `json:"studentUserIds"`
}
