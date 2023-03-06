package request

type BlockTypeAdd struct {
	Name string `json:"blockTypeName"`
}

type BlockTypeUpdate struct {
	ID   string `json:"id"`
	Name string `json:"blockTypeName"`
}
