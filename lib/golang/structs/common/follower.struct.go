package common

type Followers struct {
	Followers []Follower `json:"followers"`
}

type Follower struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
