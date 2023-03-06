package common

type SlimBlock struct {
	ID   string   `json:"id"`
	Name string   `json:"blockName"`
	Key  *SlimKey `json:"key"`
}
