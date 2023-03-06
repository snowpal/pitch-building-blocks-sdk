package request

type PodAdd struct {
	Name string `json:"podName"`
}

type PodUpdate struct {
	ID                string   `json:"id"`
	Name              string   `json:"podName"`
	Description       string   `json:"podDescription"`
	TaggedUserIds     []string `json:"taggedUserIds"`
	SimpleDescription string   `json:"simpleDescription"`
	DueDate           string   `json:"podDueDate"`
	Color             string   `json:"podColor"`
	Tags              string   `json:"blockTags"`
	KanbanMode        bool     `json:"kanbanMode"`
}

type PodBulkUpdate struct {
	PodIds []string `json:"podIds"`
}
