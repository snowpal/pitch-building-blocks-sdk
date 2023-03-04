package response

type ResourceAttribute struct {
	Name    string `json:"attributeName"`
	CanHide bool   `json:"canHide"`
}

type ResourceAttributes struct {
	KeyAttributes   []ResourceAttribute
	BlockAttributes []ResourceAttribute
	PodAttributes   []ResourceAttribute
}
