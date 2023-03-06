package request

type KeyAdd struct {
	Name string `json:"keyName"`
	Type string `json:"keyType"`
}

type KeyUpdate struct {
	ID                string `json:"id"`
	Name              string `json:"keyName"`
	Description       string `json:"keyDescription"`
	SimpleDescription string `json:"simpleDescription"`
	Color             string `json:"color"`
	Tags              string `json:"tags"`
	KanbanMode        bool   `json:"kanbanMode"`
}
