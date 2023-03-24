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

type UnreadCount struct {
	UnreadCount int `json:"unreadCount"`
}

type AllowArchivalReqBody struct {
	AllowArchival bool `json:"allowArchival"`
}

type ResourceIdParam struct {
	KeyId   string
	BlockId string
	PodId   string
}

type SearchUsersParam struct {
	SearchToken string
	ResourceIds ResourceIdParam
}

type AclParam struct {
	UserId      string
	ResourceIds ResourceIdParam
}
