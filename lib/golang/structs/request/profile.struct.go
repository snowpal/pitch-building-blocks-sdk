package request

type Profile struct {
	FirstName   *string `json:"firstName"`
	MiddleName  *string `json:"middleName"`
	LastName    *string `json:"lastName"`
	PhoneNumber *string `json:"phoneNumber"`
}
