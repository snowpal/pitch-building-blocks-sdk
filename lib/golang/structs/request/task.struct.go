package request

type TaskAdd struct {
	Name string `json:"taskName"`
}

type TaskUpdate struct {
	ID        string `json:"id"`
	Name      string `json:"taskName"`
	DueDate   string `json:"dueDate"`
	Completed bool   `json:"isCompleted"`
}

type TasksReorder struct {
	TaskIds []string `json:"taskIds"`
}
