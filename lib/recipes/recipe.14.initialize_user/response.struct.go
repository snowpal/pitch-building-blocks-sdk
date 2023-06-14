package recipes

import "github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"

type KeyWithContent struct {
	Key    response.Key
	Blocks []BlockWithContent
	Pods   []response.Pod
}

type BlockWithContent struct {
	Block     response.Block
	BlockPods []response.Pod
}

type AllKeys struct {
	CustomKey    KeyWithContent
	ClassroomKey KeyWithContent
	ProjectKey   KeyWithContent
}
