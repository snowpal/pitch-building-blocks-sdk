package request

type PodTypeAdd struct {
	Name string `json:"podTypeName"`
}

type PodTypeUpdate struct {
	ID   string `json:"id"`
	Name string `json:"podTypeName"`
}
