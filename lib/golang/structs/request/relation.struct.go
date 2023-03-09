package request

type SearchRelationWithParam struct {
	Token          string  `json:"token"`
	KeyId          *string `json:"keyId"`
	BlockId        *string `json:"blockId"`
	CurrentKeyId   *string `json:"currentKeyId"`
	CurrentBlockId *string `json:"currentBlockId"`
	CurrentPodId   *string `json:"currentPodId"`
}

type RelationWithParam struct {
	SourceKeyId   *string `json:"SourceKeyId"`
	SourceBlockId *string `json:"SourceBlockId"`
	TargetKeyId   *string `json:"TargetKeyId"`
	TargetBlockId *string `json:"TargetBlockId"`
}
