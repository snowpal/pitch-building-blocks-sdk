package response

import "development/go/recipes/lib/building-blocks/structs/common"

type StudentGradeForBlockAndPod struct {
	ID           string                `json:"id"`
	Name         string                `json:"blockName"`
	Key          common.SlimKey        `json:"key"`
	Pod          *common.SlimPod       `json:"pod"`
	StudentGrade *StudentGrade         `json:"scaleValue"`
	Pods         *[]StudentGradeForPod `json:"pods"`
	Students     *[]Student            `json:"students"`
}

type StudentGradeForPod struct {
	ID           string        `json:"id"`
	Name         string        `json:"podName"`
	StudentGrade *StudentGrade `json:"scaleValue"`
}

type StudentGrade struct {
	ScaleValue   string `json:"scaleValue"`
	NumericScale int    `json:"numericScale"`
	Published    bool   `json:"published"`
	PublishedOn  string `json:"publishedOn"`
}

type Students struct {
	Students []Student `json:"students"`
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
	StudentGrade  *StudentGrade     `json:"scaleValue"`
	Key           *common.SlimKey   `json:"key"`
	Block         *common.SlimBlock `json:"block"`
}
