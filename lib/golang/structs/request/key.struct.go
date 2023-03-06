package request

type Keys struct {
	BatchIndex *int `json:"batchIndex"`
}

type Key struct {
	ID                *string `json:"id"`
	Name              *string `json:"keyName"`
	Type              *string `json:"keyType"`
	Description       *string `json:"keyDescription"`
	SimpleDescription *string `json:"simpleDescription"`
	Color             *string `json:"color"`
	Tags              *string `json:"tags"`
	KanbanMode        *bool   `json:"kanbanMode"`
}

type KeyByTemplate struct {
	TemplateId   string `json:"templateId"`
	ExcludeTasks *bool  `json:"excludeTasks"`
}

type KeyWithParam struct {
	KeyType *string   `json:"keyType"`
	KeyIds  *[]string `json:"keyIds"`
}
