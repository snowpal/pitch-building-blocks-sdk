package common

type SlimPod struct {
	ID     string       `json:"id"`
	Name   string       `json:"podName"`
	Key    *SlimKey     `json:"key"`
	Blocks *[]SlimBlock `json:"blocks"`
}
