package request

type Scales struct {
	IncludeCounts *bool `json:"includeCounts"`
	ExcludeEmpty  *bool `json:"excludeEmpty"`
}

type Scale struct {
	ID          *string   `json:"id"`
	Name        *string   `json:"scaleName"`
	Type        *string   `json:"scaleType"`
	ScaleValues *[]string `json:"scaleValues"`
}
