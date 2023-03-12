package response

import "development/go/recipes/lib/golang/structs/common"

type Comments struct {
	Comments []Comment `json:"comments"`
}

type Comment struct {
	ID          string       `json:"id"`
	CommentText string       `json:"commentText"`
	TaggedUsers []TaggedUser `json:"taggedUsers"`

	CanEdit   *bool `json:"canEdit"`
	CanDelete *bool `json:"canDelete"`

	Key   *common.SlimKey   `json:"key"`
	Block *common.SlimBlock `json:"block"`
	Pod   *common.SlimPod   `json:"pod"`

	Creator      common.ResourceCreator  `json:"creator"`
	Modifier     common.ResourceModifier `json:"modifier"`
	LastModified string                  `json:"lastModified"`
}

type RecentComments struct {
	Comments []RecentComment `json:"comments"`
}

type RecentComment struct {
	ID          string `json:"id"`
	CommentText string `json:"commentText"`

	Key    *common.SlimKey     `json:"key"`
	Block  *common.SlimBlock   `json:"block"`
	Blocks *[]common.SlimBlock `json:"blocks"`
	Pod    *common.SlimPod     `json:"pod"`

	StudentId *string `json:"studentId"`

	Creator      common.ResourceCreator  `json:"creator"`
	Modifier     common.ResourceModifier `json:"modifier"`
	LastModified string                  `json:"lastModified"`
}
