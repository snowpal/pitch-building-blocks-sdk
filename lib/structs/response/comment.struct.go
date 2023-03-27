package response

import (
	common2 "github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
)

type Comments struct {
	Comments []Comment `json:"comments"`
}

type Comment struct {
	ID          string       `json:"id"`
	CommentText string       `json:"commentText"`
	TaggedUsers []TaggedUser `json:"taggedUsers"`

	CanEdit   *bool `json:"canEdit"`
	CanDelete *bool `json:"canDelete"`

	Key   *common2.SlimKey   `json:"key"`
	Block *common2.SlimBlock `json:"block"`
	Pod   *common2.SlimPod   `json:"pod"`

	Creator      common2.ResourceCreator  `json:"creator"`
	Modifier     common2.ResourceModifier `json:"modifier"`
	LastModified string                   `json:"lastModified"`
}

type RecentComments struct {
	Comments []RecentComment `json:"comments"`
}

type RecentComment struct {
	ID          string `json:"id"`
	CommentText string `json:"commentText"`

	Key    *common2.SlimKey     `json:"key"`
	Block  *common2.SlimBlock   `json:"block"`
	Blocks *[]common2.SlimBlock `json:"blocks"`
	Pod    *common2.SlimPod     `json:"pod"`

	StudentId *string `json:"studentId"`

	Creator      common2.ResourceCreator  `json:"creator"`
	Modifier     common2.ResourceModifier `json:"modifier"`
	LastModified string                   `json:"lastModified"`
}
