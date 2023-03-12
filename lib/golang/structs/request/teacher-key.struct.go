package request

type TeacherBlock struct {
	StudentUserIds []string `json:"studentUserIds"`
	PodIds         []string `json:"podIds"`
}
