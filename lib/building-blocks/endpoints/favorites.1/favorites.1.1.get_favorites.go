package favorites

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetFavorites(jwtToken string) ([]response.Favorite, error) {
	resFavorites := response.Favorites{}
	route, err := helpers.GetRoute(building_blocks.RouteFavoritesGetFavorites)
	if err != nil {
		fmt.Println(err)
		return resFavorites.Favorites, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resFavorites.Favorites, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resFavorites.Favorites, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resFavorites.Favorites, err
	}

	err = json.Unmarshal(body, &resFavorites)
	if err != nil {
		fmt.Println(err)
		return resFavorites.Favorites, err
	}
	return resFavorites.Favorites, nil
}
