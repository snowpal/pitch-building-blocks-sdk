package recipes

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/endpoints/favorites"
	recipes2 "development/go/recipes/lib/helpers/recipes"
	"development/go/recipes/lib/structs/common"
	response2 "development/go/recipes/lib/structs/response"
	log "github.com/sirupsen/logrus"
)

const (
	FavKeyName   = "FavoriteKey"
	FavBlockName = "FavoriteBlock"
)

func AddFavorite() {
	log.Info("Objective: Add and Remove Favorites")
	_, err := recipes2.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes2.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	log.Info("Create a key and a block into it. Then add that block as favorite")
	var favorite response2.AddFavorite
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

func removeFavorite(user response2.User, favorite response2.AddFavorite) error {
	err := favorites.DeleteFavorite(user.JwtToken, favorite.ID)
	if err != nil {
		return err
	}
	return nil
}

func addFavorite(user response2.User) (response2.AddFavorite, error) {
	var favorite response2.AddFavorite
	key, err := recipes2.AddCustomKey(user, FavKeyName)
	if err != nil {
		return favorite, err
	}
	block, err := recipes2.AddBlock(user, FavBlockName, key)
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
