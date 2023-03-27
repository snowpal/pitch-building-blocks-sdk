package response

import (
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
)

type Favorites struct {
	Favorites []Favorite `json:"favorites"`
}

type AddFavorite struct {
	ID string `json:"id"`
}

type Favorite struct {
	ID       string                  `json:"id"`
	Modifier common.ResourceModifier `json:"modifier"`
	Resource FavoriteResource        `json:"resource"`
}

type FavoriteKey struct {
	ID       string                  `json:"id"`
	Name     string                  `json:"keyName"`
	Modifier common.ResourceModifier `json:"modifier"`
}

type FavoriteBlock struct {
	ID       string                  `json:"id"`
	Name     string                  `json:"blockName"`
	Modifier common.ResourceModifier `json:"modifier"`
}

type FavoriteResource struct {
	ID           string                  `json:"id"`
	ResourceType string                  `json:"resourceType"`
	KeyName      *string                 `json:"keyName"`
	BlockName    *string                 `json:"blockName"`
	PodName      *string                 `json:"podName"`
	Key          *FavoriteKey            `json:"key"`
	Block        *FavoriteBlock          `json:"block"`
	Modifier     common.ResourceModifier `json:"modifier"`
}
