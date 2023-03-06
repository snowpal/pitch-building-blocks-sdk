package request

type EventsWithParam struct {
	StartDate string  `json:"startDate"`
	EndDate   *string `json:"endDate"`
}

type StandaloneEvent struct {
	Description *string `json:"description"`
	StartTime   *string `json:"eventStartTime"`
	EndTime     *string `json:"eventEndTime"`
}
