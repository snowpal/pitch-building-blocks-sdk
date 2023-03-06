package response

import "development/go/recipes/lib/golang/structs/common"

type Comments struct {
	Comments []Comment `json:"comments"`
}

type Comment struct {
	ID           string                  `json:"id"`
	CommentText  string                  `json:"commentText"`
	Key          *common.SlimKey         `json:"key"`
	Block        *common.SlimBlock       `json:"block"`
	Creator      common.ResourceCreator  `json:"creator"`
	Modifier     common.ResourceModifier `json:"modifier"`
	LastModified string                  `json:"lastModified"`
	CanEdit      *bool                   `json:"canEdit"`
	CanDelete    *bool                   `json:"canDelete"`
	TaggedUsers  []User                  `json:"taggedUsers"`
}

type RecentComments struct {
	Comments []RecentComment `json:"comments"`
}

type RecentComment struct {
	ID           string                  `json:"id"`
	CommentText  string                  `json:"commentText"`
	StudentId    string                  `json:"studentId"`
	Key          *common.SlimKey         `json:"key"`
	Block        *common.SlimBlock       `json:"block"`
	Blocks       *[]common.SlimBlock     `json:"blocks"`
	Pod          *common.SlimPod         `json:"pod"`
	Creator      common.ResourceCreator  `json:"creator"`
	Modifier     common.ResourceModifier `json:"modifier"`
	LastModified string                  `json:"lastModified"`
}
