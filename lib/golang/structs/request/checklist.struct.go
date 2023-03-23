package request

type ChecklistIdParam struct {
	KeyId           string
	BlockId         *string
	PodId           *string
	ChecklistId     *string
	ChecklistItemId *string
}

type ChecklistReqBody struct {
	Title string `json:"title"`
}

type ReorderChecklistsReqBody struct {
	ChecklistIds string `json:"checklistIds"`
}

type ChecklistItemReqBody struct {
	Title         string  `json:"title"`
	Completed     *bool   `json:"completed"`
	TaggedUserIds *string `json:"taggedUserIds"`
}

type ReorderChecklistItemsReqBody struct {
	ChecklistItemIds string `json:"checklistItemIds"`
}
