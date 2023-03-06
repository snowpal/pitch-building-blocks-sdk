package request

type ProjectList struct {
	ID             string   `json:"id"`
	Name           string   `json:"projectListName"`
	ProjectListIds []string `json:"projectListIds"`
	UserIds        []string `json:"userIds"`
}

type ProjectListWithParam struct {
	KeyId          string `json:"keyId"`
	TargetKeyId    string `json:"targetKeyId"`
	TargetBlockId  string `json:"targetBlockIdId"`
	TargetPosition string `json:"targetPosition"`
}

type ProjectPodWithParam struct {
	KeyId               string `json:"keyId"`
	BlockId             string `json:"blockId"`
	TargetKeyId         string `json:"targetKeyId"`
	TargetBlockId       string `json:"targetBlockIdId"`
	TargetProjectListId string `json:"targetProjectListId"`

	PodIds   *[]string `json:"podIds"`
	AllPods  *bool     `json:"allPods"`
	AllTasks *bool     `json:"allTasks"`
}

type ProjectPodByTemplate struct {
	KeyId         string `json:"keyId"`
	ProjectListId string `json:"projectListId"`
	TemplateId    string `json:"templateId"`
	ExcludeTasks  *bool  `json:"excludeTasks"`
}
