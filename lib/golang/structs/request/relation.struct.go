package request

type KeyToKeyRelationParam struct {
	KeyId       string
	TargetKeyId string
}

type KeyToBlockRelationParam struct {
	KeyId         string
	TargetBlockId string
}

type KeyToPodRelationParam struct {
	KeyId string

	TargetPodId string
	TargetKeyId string
}

type KeyToBlockPodRelationParam struct {
	KeyId string

	TargetPodId   string
	TargetKeyId   string
	TargetBlockId string
}

type BlockToBlockRelationParam struct {
	BlockId       string
	TargetBlockId string
}

type BlockToPodRelationParam struct {
	BlockId string

	TargetPodId string
	TargetKeyId string
}

type BlockToBlockPodRelationParam struct {
	BlockId string

	TargetPodId   string
	TargetKeyId   string
	TargetBlockId string
}

type PodToPodRelationParam struct {
	PodId       string
	SourceKeyId string

	TargetPodId string
	TargetKeyId string
}

type PodToBlockPodRelationParam struct {
	PodId       string
	SourceKeyId string

	TargetPodId   string
	TargetKeyId   string
	TargetBlockId string
}

type BlockPodToBlockPodRelationParam struct {
	PodId         string
	SourceKeyId   string
	SourceBlockId string

	TargetPodId   string
	TargetKeyId   string
	TargetBlockId string
}
