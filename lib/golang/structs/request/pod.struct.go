package request

type AddPodReqBody struct {
	Name string `json:"podName"`
}

type GetBlockPodsParam struct {
	KeyId      string
	BlockId    string
	BatchIndex int
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

	AllowArchival *bool `json:"allowArchival"`
	Completed     *bool `json:"podCompleted"`
	KanbanMode    *bool `json:"kanbanMode"`

	TaggedUserIds *[]string `json:"taggedUserIds"`
	PodIds        *[]string `json:"podIds"`
}

type CopyMovePodParam struct {
	PodId       string `json:"podId"`
	KeyId       string `json:"keyId"`
	TargetKeyId string `json:"targetKeyId"`

	BlockId       *string `json:"blockId"`
	TargetBlockId *string `json:"targetBlockId"`

	AllTasks      *bool `json:"allTasks"`
	AllChecklists *bool `json:"allChecklists"`
}

type SharePod struct {
	Acl            string    `json:"podAcl"`
	PodIds         *[]string `json:"podIds"`
	StudentUserIds *[]string `json:"studentUserIds"`
}
