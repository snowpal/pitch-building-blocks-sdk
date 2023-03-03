package response

type Key struct {
	Name              string `json:"keyName"`
	Description       string `json:"keyDescription"`
	SimpleDescription string `json:"simpleDescription"`
	Color             string `json:"color"`
	Tags              string `json:"tags"`
	KanbanMode        bool   `json:"kanbanMode"`
}
