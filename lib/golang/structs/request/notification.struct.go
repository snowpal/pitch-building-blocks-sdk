package request

type Notification struct {
	Unread          *bool     `json:"unread"`
	NotificationIds *[]string `json:"notificationIds"`
}
