package request

type ScaleAdd struct {
	Name        string   `json:"scaleName"`
	Type        string   `json:"scaleType"`
	ScaleValues []string `json:"scaleValues"`
}

type ScaleUpdate struct {
	ID          string   `json:"id"`
	Name        string   `json:"scaleName"`
	Type        string   `json:"scaleType"`
	ScaleValues []string `json:"scaleValues"`
}
