package response

import "development/go/recipes/lib/golang/structs/common"

type SlimKeysExt struct {
	Keys []SlimKeyExt `json:"keys"`
}

type SlimKeyExt struct {
	ID           string           `json:"id"`
	Name         string           `json:"keyName"`
	Type         string           `json:"keyType"`
	Blocks       []SlimBlockExt   `json:"blocks"`
	Pods         []common.SlimPod `json:"pods"`
	LastModified string           `json:"lastModified"`
}

type SlimBlockExt struct {
	ID           string           `json:"id"`
	Name         string           `json:"keyName"`
	Pods         []common.SlimPod `json:"pods"`
	LastModified string           `json:"lastModified"`
}

type FilteredKeys struct {
	Keys []FilteredKey `json:"keys"`
}

type BlocksAndPods struct {
	Blocks []common.SlimBlock `json:"blocks"`
	Pods   []common.SlimPod   `json:"pods"`
}

type FilteredKey struct {
	ID               string        `json:"id"`
	Name             string        `json:"keyName"`
	Type             string        `json:"keyType"`
	CreatedByMe      BlocksAndPods `json:"createdByMe"`
	SharedWithMe     BlocksAndPods `json:"sharedWithMe"`
	SharedWithOthers BlocksAndPods `json:"sharedWithOthers"`
}

type BlockTypesKeys struct {
	Keys []BlockTypesKey `json:"keys"`
}

type BlockTypesKey struct {
	ID           string       `json:"id"`
	Name         string       `json:"keyName"`
	Type         string       `json:"keyType"`
	BlockTypes   *[]BlockType `json:"blockTypes"`
	LastModified string       `json:"lastModified"`
}

type PodTypesKeysKeyPod struct {
	KeyPod PodTypesKeys `json:"keyPod"`
}

type PodTypesKeysBlockPod struct {
	BlockPod PodTypesKeys `json:"blockPod"`
}

type PodTypesKeysOtherPod struct {
	OtherPod PodTypesKeys `json:"otherPod"`
}

type PodTypesKeys struct {
	Keys *[]PodTypesKey `json:"keys"`
	Key  *PodTypesKey   `json:"key"`
}

type PodTypesKey struct {
	ID           string     `json:"id"`
	Name         string     `json:"keyName"`
	Type         string     `json:"keyType"`
	PodTypes     *[]PodType `json:"podTypes"`
	LastModified string     `json:"lastModified"`
}

type ScalesKeysBlock struct {
	Block ScalesKeys `json:"block"`
}

type ScalesKeysKeyPod struct {
	KeyPod ScalesKeys `json:"keyPod"`
}

type ScalesKeysBlockPod struct {
	BlockPod ScalesKeys `json:"blockPod"`
}

type ScalesKeysOtherPod struct {
	OtherPod ScalesKeys `json:"otherPod"`
}

type ScalesKeys struct {
	Keys *[]ScalesKey `json:"keys"`
	Key  *ScalesKey   `json:"key"`
}

type ScalesKey struct {
	ID           string   `json:"id"`
	Name         string   `json:"keyName"`
	Type         string   `json:"keyType"`
	Scales       *[]Scale `json:"scales"`
	LastModified string   `json:"lastModified"`
}

type TaskStatus struct {
	Complete   int `json:"complete"`
	Incomplete int `json:"incomplete"`
}

type TasksStatusKeys struct {
	Keys []TasksStatusKey `json:"keys"`
}

type TasksStatusKey struct {
	ID           string     `json:"id"`
	Name         string     `json:"keyName"`
	TaskStatus   TaskStatus `json:"taskStatus"`
	LastModified string     `json:"lastModified"`
}

type TasksStatusBlock struct {
	ID           string         `json:"id"`
	Name         string         `json:"blockName"`
	TaskStatus   TaskStatus     `json:"taskStatus"`
	Key          common.SlimKey `json:"key"`
	LastModified string         `json:"lastModified"`
}

type LinkedResourcesKey struct {
	SlimKey common.SlimKey `json:"key"`
}

type LinkedResources struct {
	CurrentKey LinkedResourcesKey `json:"currentKey"`
	SharedKey  LinkedResourcesKey `json:"sharedKey"`
	Keys       *[]SlimKeyExt      `json:"keys"`
	Blocks     []SlimBlockExt     `json:"blocks"`
}

type BlockScaleValue struct {
	ID           string `json:"id"`
	Name         string `json:"blockName"`
	ScaleValue   string `json:"scaleValue"`
	NumericScale string `json:"numericScale"`
}

type PodScaleValue struct {
	ID           string `json:"id"`
	Name         string `json:"podName"`
	ScaleValue   string `json:"scaleValue"`
	NumericScale string `json:"numericScale"`
}

type ScaleValues struct {
	Scale  Scale              `json:"scale"`
	Key    common.SlimKey     `json:"key"`
	Block  *common.SlimBlock  `json:"block"`
	Blocks *[]BlockScaleValue `json:"blocks"`
	Pods   []PodScaleValue    `json:"pods"`
}

type BlockGrade struct {
	ID       string         `json:"id"`
	Name     string         `json:"blockName"`
	Key      common.SlimKey `json:"key"`
	Students []Student      `json:"students"`
}

type PodGrade struct {
	ID       string           `json:"id"`
	Name     string           `json:"podName"`
	Key      common.SlimKey   `json:"key"`
	Block    common.SlimBlock `json:"block"`
	Students []Student        `json:"students"`
}
