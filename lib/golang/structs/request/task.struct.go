package request

type AddTaskReqBody struct {
	Name string `json:"taskName"`
}

type UpdateTaskReqBody struct {
	Name      *string `json:"taskName"`
	DueDate   *string `json:"dueDate"`
	Completed *bool   `json:"isCompleted"`
}

type AssignTaskReqBody struct {
	UserIds string `json:"userIds"`
}

type ReorderTasksReqBody struct {
	TaskIds string `json:"taskIds"`
}

type TaskIdParam struct {
	KeyId   string
	BlockId *string
	PodId   *string
	TaskId  *string
}
