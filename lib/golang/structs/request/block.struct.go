package request

type AddBlockReqBody struct {
	Name string `json:"blockName"`
}

type GetBlocksParam struct {
	KeyId            string
	BatchIndex       int
	WriteOrHigherAcl bool
	Filter           *string
}

type CopyMoveBlockParam struct {
	BlockId     string
	KeyId       string
	TargetKeyId string

	PodIds        *[]string
	AllPods       *bool
	AllTasks      *bool
	AllChecklists *bool
}

type ShareBlock struct {
	Acl      string    `json:"blockAcl"`
	PodIds   *[]string `json:"podIds"`
	BlockIds *[]string `json:"blockIds"`
}
