package request

type Task struct {
	ID        *string   `json:"id"`
	Name      *string   `json:"taskName"`
	DueDate   *string   `json:"dueDate"`
	Completed *bool     `json:"isCompleted"`
	TaskIds   *[]string `json:"taskIds"`
	UserIds   *[]string `json:"userIds"`
}
