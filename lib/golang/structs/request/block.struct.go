package request

type BlockAdd struct {
	Name string `json:"blockName"`
}

type BlockUpdate struct {
	ID                string   `json:"id"`
	Name              string   `json:"blockName"`
	BlockId           string   `json:"blockId"`
	Description       string   `json:"blockDescription"`
	TaggedUserIds     []string `json:"taggedUserIds"`
	SimpleDescription string   `json:"simpleDescription"`
	DueDate           string   `json:"blockDueDate"`
	StartTime         string   `json:"blockStartTime"`
	EndTime           string   `json:"blockEndTime"`
	Color             string   `json:"blockColor"`
	Tags              string   `json:"blockTags"`
	Completed         bool     `json:"blockCompleted"`
	KanbanMode        bool     `json:"kanbanMode"`
}

type BlockBulkUpdate struct {
	BlockIds []string `json:"blockIds"`
}
