package response

import "development/go/recipes/lib/golang/structs/common"

type AllEvents struct {
	DueDateEvent    *DueDateEvent    `json:"dueDate"`
	EndDateEvent    *EndDateEvent    `json:"endDate"`
	SchedulerEvents []SchedulerEvent `json:"schedulerEvents"`
}

type EndDateEvent struct {
	Blocks []BlockEvent `json:"blocks"`
}

type DueDateEvent struct {
	Tasks  TasksEvent   `json:"tasks"`
	Blocks []BlockEvent `json:"blocks"`
	Pods   []PodEvent   `json:"pods"`
}

type BlockEvent struct {
	ID          string `json:"id"`
	Name        string `json:"blockName"`
	Description string `json:"blockDescription"`

	DueDate   *string `json:"blockDueDate"`
	StartTime *string `json:"blockStartTime"`
	EndTime   *string `json:"blockEndTime"`

	Key common.SlimKey `json:"key"`
}

type PodEvent struct {
	ID      string `json:"id"`
	Name    string `json:"podName"`
	DueDate string `json:"podDueDate"`

	Key   common.SlimKey   `json:"key"`
	Block common.SlimBlock `json:"block"`
}

type TasksEvent struct {
	KeyTasks   []TaskEvent `json:"keys"`
	BlockTasks []TaskEvent `json:"blocks"`
	PodTasks   []TaskEvent `json:"pods"`
}

type TaskEvent struct {
	ID      string            `json:"id"`
	Name    string            `json:"taskName"`
	DueDate string            `json:"taskDueDate"`
	Key     common.SlimKey    `json:"key"`
	Block   *common.SlimBlock `json:"block"`
	Pod     *common.SlimPod   `json:"pod"`
}

type SchedulerEvent struct {
	SchedulerId      string            `json:"schedulerId"`
	StandaloneEvents []StandaloneEvent `json:"standaloneEvents"`
}

type StandaloneEvent struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	StartTime   string `json:"eventStartTime"`
	EndTime     string `json:"eventEndTime"`
}
