package request

type AddPodReqBody struct {
	Name string `json:"podName"`
}

type UpdatePodDescReqBody struct {
	Description   string  `json:"podDescription"`
	TaggedUserIds *string `json:"taggedUserIds"`
}

type BulkArchivePodsReqBody struct {
	PodIds string `json:"podIds"`
}

type UpdatePodStatusReqBody struct {
	Completed bool `json:"podCompleted"`
}

type UpdatePodReqBody struct {
	Name              *string `json:"podName"`
	SimpleDescription *string `json:"simpleDescription"`
	DueDate           *string `json:"podDueDate"`
	Color             *string `json:"podColor"`
	Tags              *string `json:"podTags"`
	KanbanMode        *bool   `json:"kanbanMode"`
}

type PodAclReqBody struct {
	Acl string `json:"podAcl"`
}

type PodBulkShareReqBody struct {
	Acl    string `json:"podAcl"`
	PodIds string `json:"podIds"`
}

type GetPodsParam struct {
	KeyId      string
	BlockId    *string
	BatchIndex int
}

type AddPodTypeIdParam struct {
	PodId     string
	PodTypeId string
	KeyId     string
	BlockId   *string
}

type PodByTemplateParam struct {
	KeyId        string
	BlockId      *string
	TemplateId   string
	ExcludeTasks bool
}

type CopyMovePodParam struct {
	PodId       string
	KeyId       string
	TargetKeyId string

	BlockId       string
	TargetBlockId string

	AllTasks      bool
	AllChecklists bool
}
