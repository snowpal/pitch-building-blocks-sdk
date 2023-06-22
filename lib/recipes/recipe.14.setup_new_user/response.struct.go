package setupnewuser

import "github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"

type KeyWithResources struct {
	Key    response.Key
	Blocks []BlockWithPods
	Pods   []response.Pod
}

func (key *KeyWithResources) firstBlock() BlockWithPods {
	return key.Blocks[0]
}

func (key *KeyWithResources) firstPod() response.Pod {
	return key.Pods[0]
}

func (key *KeyWithResources) firstBlockPod() response.Pod {
	return key.Blocks[0].firstPod()
}

type BlockWithPods struct {
	Block     response.Block
	BlockPods []response.Pod
}

func (block *BlockWithPods) firstPod() response.Pod {
	return block.BlockPods[0]
}

type AllKeys struct {
	CustomKey  KeyWithResources
	TeacherKey KeyWithResources
	ProjectKey KeyWithResources
}
