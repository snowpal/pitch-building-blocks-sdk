package request

type AddKeyReqBody struct {
	Type string `json:"keyType"`
	Name string `json:"keyName"`
}
