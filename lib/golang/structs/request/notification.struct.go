package request

type NotificationUpdate struct {
	Unread bool `json:"unread"`
}

type NotificationBulkUpdate struct {
	NotificationIds []string `json:"notificationIds"`
}
