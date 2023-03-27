package common

type SlimPod struct {
	ID     string      `json:"id"`
	Name   string      `json:"podName"`
	Key    *SlimKey    `json:"key"`
	Block  *SlimBlock  `json:"block"`
	Blocks []SlimBlock `json:"blocks"`
}
