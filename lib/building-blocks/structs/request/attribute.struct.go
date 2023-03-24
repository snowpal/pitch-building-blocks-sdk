package request

type ResourceAttributeReqBody struct {
	AttributeNames string `json:"attributeNames"`
	ShowAttribute  bool   `json:"showAttribute"`
}
