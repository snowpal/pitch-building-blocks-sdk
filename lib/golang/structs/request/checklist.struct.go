package request

type Checklist struct {
	ID           *string   `json:"id"`
	Title        *string   `json:"checklistTitle"`
	ChecklistIds *[]string `json:"checklistIds"`
}

type ChecklistItem struct {
	ID        *string `json:"id"`
	ItemText  *string `json:"checklistItemText"`
	Completed *bool   `json:"completed"`

	TaggedUserIds    *[]string `json:"taggedUserIds"`
	ChecklistItemIds *[]string `json:"checklistItemIds"`
}
