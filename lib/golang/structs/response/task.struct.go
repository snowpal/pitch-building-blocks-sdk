package response

import "development/go/recipes/lib/golang/structs/common"

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	ID        string                  `json:"id"`
	Name      string                  `json:"taskName"`
	DueDate   string                  `json:"taskDueDate"`
	Completed bool                    `json:"isCompleted"`
	Assignees []TaggedUser            `json:"assignees"`
	Key       *common.SlimKey         `json:"key"`
	Block     *common.SlimBlock       `json:"block"`
	Creator   common.ResourceCreator  `json:"creator"`
	Modifier  common.ResourceModifier `json:"modifier"`
}
