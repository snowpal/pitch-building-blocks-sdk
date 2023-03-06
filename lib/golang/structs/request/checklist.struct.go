package request

type ChecklistAdd struct {
	Title string `json:"checklistTitle"`
}

type ChecklistUpdate struct {
	ID    string `json:"id"`
	Title string `json:"checklistTitle"`
}

type ChecklistItemAdd struct {
	ItemText      string   `json:"checklistItemText"`
	TaggedUserIds []string `json:"taggedUserIds"`
}

type ChecklistItemUpdate struct {
	ID            string   `json:"id"`
	ItemText      string   `json:"checklistItemText"`
	TaggedUserIds []string `json:"taggedUserIds"`
	Completed     bool     `json:"completed"`
}

type ChecklistsReorder struct {
	ChecklistIds []string `json:"checklistIds"`
}

type ChecklistItemsReorder struct {
	ChecklistItemIds []string `json:"checklistItemIds"`
}
