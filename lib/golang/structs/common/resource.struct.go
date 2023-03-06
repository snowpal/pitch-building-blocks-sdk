package common

type DisplayAttribute struct {
	Name string `json:"attributeName"`
	Show bool   `json:"show"`
}

type ResourceCreator struct {
	CreatedOn string `json:"createdOn"`
	CreatedBy string `json:"createdBy"`
}

type ResourceModifier struct {
	ModifiedOn string `json:"modifiedOn"`
	ModifiedBy string `json:"modifiedBy"`
}

type AllowArchival struct {
	AllowArchival bool `json:"allowArchival"`
}
