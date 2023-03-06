package response

import "development/go/recipes/lib/golang/structs/common"

type Notifications struct {
	Notifications []Notification `json:"notifications"`
}

type Notification struct {
	ID       string               `json:"id"`
	Type     string               `json:"badgeType"`
	Text     string               `json:"badgeText"`
	Unread   bool                 `json:"unread"`
	Resource NotificationResource `json:"resource"`
}

type NotificationResource struct {
	ID           string          `json:"id"`
	ResourceType string          `json:"resourceType"`
	IsTask       *bool           `json:"isTask"`
	KeyName      *string         `json:"keyName"`
	BlockName    *string         `json:"blockName"`
	PodName      *string         `json:"podName"`
	Key          *common.SlimKey `json:"key"`
	StartTime    *string         `json:"eventStartTime"`
	EndTime      *string         `json:"eventEndTime"`
}

type Performer struct {
	Username  string `json:"userName"`
	FirstName string `json:"firstName"`
}

type NotificationEvent struct {
	PerformedOn string     `json:"performedOn"`
	PerformedBy *Performer `json:"performedBy"`
}
