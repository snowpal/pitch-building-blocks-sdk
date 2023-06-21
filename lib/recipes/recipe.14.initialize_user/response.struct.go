package recipes

import "github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"

type KeyWithResources struct {
	Key    response.Key
	Blocks []BlockWithPods
	Pods   []response.Pod
}

type BlockWithPods struct {
	Block     response.Block
	BlockPods []response.Pod
}

type AllKeys struct {
	CustomKey  KeyWithResources
	TeacherKey KeyWithResources
	ProjectKey KeyWithResources
}
