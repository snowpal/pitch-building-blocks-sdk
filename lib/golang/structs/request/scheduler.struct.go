package request

type EventDateParam struct {
	StartDate string
	EndDate   *string
}

type StandaloneEventReqBody struct {
	Description *string `json:"description"`
	StartTime   *string `json:"eventStartTime"`
	EndTime     *string `json:"eventEndTime"`
}
