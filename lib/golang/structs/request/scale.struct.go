package response

type Scale struct {
	Name        string   `json:"scaleName"`
	Type        string   `json:"scaleType"`
	ScaleValues []string `json:"scaleValues"`
}
