package response

import "development/go/recipes/lib/golang/structs/common"

type Dashboard struct {
	Dashboard DashboardResources `json:"dashboard"`
}

type DashboardResources struct {
	RecentlyModifiedResources *RecentlyModified `json:"recentlyModified"`
	ShortlyDueResources       *DueShortly       `json:"dueShortly"`
}

type RecentlyModified struct {
	Keys   *[]common.SlimKey `json:"keys"`
	Blocks *[]DashboardBlock `json:"blocks"`
	Pods   *[]DashboardPod   `json:"pods"`
}

type DueShortly struct {
	Blocks *[]DashboardBlock `json:"blocks"`
	Pods   *[]DashboardPod   `json:"pods"`
	Tasks  *[]DashboardTask  `json:"tasks"`
}

type DashboardBlock struct {
	ID      string `json:"id"`
	Name    string `json:"blockName"`
	DueDate string `json:"blockDueDate"`

	Key       *common.SlimKey `json:"key"`
	BlockType *BlockType      `json:"blockType"`

	Creator  common.ResourceCreator  `json:"creator"`
	Modifier common.ResourceModifier `json:"modifier"`
}

type DashboardPod struct {
	ID      string `json:"id"`
	Name    string `json:"podName"`
	DueDate string `json:"podDueDate"`

	Key    *common.SlimKey     `json:"key"`
	Blocks *[]BlockWithPodType `json:"blocks"`

	Creator  common.ResourceCreator  `json:"creator"`
	Modifier common.ResourceModifier `json:"modifier"`
}

type BlockWithPodType struct {
	ID      string          `json:"id"`
	Name    string          `json:"blockName"`
	Key     *common.SlimKey `json:"key"`
	PodType *PodType        `json:"podType"`
}

type DashboardTask struct {
	ID      string `json:"id"`
	Name    string `json:"taskName"`
	DueDate string `json:"taskDueDate"`

	Key    *common.SlimKey     `json:"key"`
	Block  *common.SlimBlock   `json:"block"`
	Pod    *common.SlimPod     `json:"pod"`
	Blocks *[]common.SlimBlock `json:"blocks"`

	Creator  common.ResourceCreator  `json:"creator"`
	Modifier common.ResourceModifier `json:"modifier"`
}

type DashboardUnreadCount struct {
	DueTasks  int `json:"dueTasks"`
	DueBlocks int `json:"dueBlocks"`
	DuePods   int `json:"duePods"`

	Notifications int `json:"notifications"`
	Conversations int `json:"conversations"`
}
