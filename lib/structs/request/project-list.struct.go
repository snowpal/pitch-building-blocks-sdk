package request

type AddProjectListReqBody struct {
	Name string `json:"projectListName"`
}

type AssignProjectPodReqBody struct {
	UserIds string `json:"userIds"`
}

type ProjectListIdParam struct {
	KeyId         string
	BlockId       string
	ProjectListId string
	PodId         *string
}

type CopyMoveProjectListParam struct {
	KeyId         string
	BlockId       string
	ProjectListId string

	TargetKeyId    string
	TargetBlockId  string
	TargetPosition int
}

type CopyMoveProjectPodParam struct {
	PodId   string
	BlockId string
	KeyId   string

	TargetKeyId         string
	TargetBlockId       string
	TargetProjectListId string
}

type CopyMoveProjectListPodsParam struct {
	KeyId         string
	BlockId       string
	ProjectListId string

	TargetKeyId         string
	TargetBlockId       string
	TargetProjectListId string

	AllTasks bool
	AllPods  bool

	PodIds []string
}
