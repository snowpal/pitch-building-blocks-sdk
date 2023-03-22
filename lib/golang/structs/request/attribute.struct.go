package request

type ResourceAttribute struct {
	AttributeNames []string `json:"attributeNames"`
	ShowAttribute  bool     `json:"showAttribute"`
}
