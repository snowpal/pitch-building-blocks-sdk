package response

type Block struct {
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
	Completed         *bool    `json:"blockCompleted"`
	KanbanMode        bool     `json:"kanbanMode"`
}
