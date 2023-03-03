package structs

type HideableAttribute struct {
	Name    string `json:"attributeName"`
	CanHide bool   `json:"canHide"`
}

type ResourceAttributes struct {
	KeyAttributes   []HideableAttribute
	BlockAttributes []HideableAttribute
	PodAttributes   []HideableAttribute
}
