package response

type Pod struct {
	Name              string   `json:"podName"`
	Description       string   `json:"podDescription"`
	TaggedUserIds     []string `json:"taggedUserIds"`
	SimpleDescription string   `json:"simpleDescription"`
	DueDate           string   `json:"podDueDate"`
	Color             string   `json:"podColor"`
	Tags              string   `json:"blockTags"`
	KanbanMode        bool     `json:"kanbanMode"`
}
