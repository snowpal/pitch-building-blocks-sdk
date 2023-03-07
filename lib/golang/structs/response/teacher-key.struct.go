package response

import "development/go/recipes/lib/golang/structs/common"

type TeacherBlock struct {
	ID         string         `json:"id"`
	Name       string         `json:"blockName"`
	Key        common.SlimKey `json:"key"`
	ScaleValue *ScaleValue    `json:"scaleValue"`
	Pods       *[]TeacherPod  `json:"pods"`
	Students   *[]Student     `json:"students"`
}

type TeacherPod struct {
	ID         string      `json:"id"`
	Name       string      `json:"podName"`
	ScaleValue *ScaleValue `json:"scaleValue"`
}

type ScaleValue struct {
	ScaleValue   string `json:"scaleValue"`
	NumericScale int    `json:"numericScale"`
	Published    bool   `json:"published"`
	PublishedOn  string `json:"publishedOn"`
}

type Student struct {
	ID            string            `json:"id"`
	ProfileID     string            `json:"profileId"`
	Email         string            `json:"email"`
	Username      string            `json:"username"`
	FirstName     string            `json:"firstName"`
	MiddleName    string            `json:"middleName"`
	LastName      string            `json:"lastName"`
	PhoneNumber   string            `json:"phoneNumber"`
	AddressUserBy string            `json:"addressUserBy"`
	UserInitial   string            `json:"userInitial"`
	AvatarName    string            `json:"avatarName"`
	AvatarUrl     string            `json:"avatarUrl"`
	BlockName     string            `json:"blockName"`
	ScaleValue    *ScaleValue       `json:"scaleValue"`
	Key           *common.SlimKey   `json:"key"`
	Block         *common.SlimBlock `json:"block"`
}
