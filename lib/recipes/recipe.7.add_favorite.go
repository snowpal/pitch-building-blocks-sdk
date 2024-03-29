package recipes

import (
	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/blocks/blocks.1"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/favorites"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/keys/keys.1"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"

	log "github.com/sirupsen/logrus"
	recipes "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers/recipes"
	response "github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

const (
	FavKeyName   = "FavoriteKey"
	FavBlockName = "FavoriteBlock"
)

func AddFavorite() {
	log.Info("Objective: Add and Remove Favorites")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	log.Info("Create a key and a block into it. Then add that block as favorite")
	var favorite response.AddFavorite
	favorite, err = addFavorite(user)
	if err != nil {
		return
	}
	log.Info(".Block added as favorite")

	err = removeFavorite(user, favorite)
	if err != nil {
		return
	}
	log.Info(".Block removed from favorite")
}

func removeFavorite(user response.User, favorite response.AddFavorite) error {
	err := favorites.DeleteFavorite(user.JwtToken, favorite.ID)
	if err != nil {
		return err
	}
	return nil
}

func addFavorite(user response.User) (response.AddFavorite, error) {
	var favorite response.AddFavorite
	key, err := keys.AddKey(
		user.JwtToken,
		request.AddKeyReqBody{
			Name: FavKeyName,
			Type: lib.KeyTypes[lib.Custom],
		})
	if err != nil {
		return favorite, err
	}
	block, err := blocks.AddBlock(
		user.JwtToken,
		request.AddBlockReqBody{Name: FavBlockName},
		key.ID)
	if err != nil {
		return favorite, err
	}
	favorite, err = favorites.AddBlockAsFavorite(
		user.JwtToken,
		common.ResourceIdParam{BlockId: block.ID, KeyId: key.ID})
	if err != nil {
		return favorite, err
	}
	return favorite, nil
}
