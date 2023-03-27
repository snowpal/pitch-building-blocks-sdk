package favorites

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetFavorites(jwtToken string) ([]response.Favorite, error) {
	resFavorites := response.Favorites{}
	route, err := helpers2.GetRoute(lib.RouteFavoritesGetFavorites)
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

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resFavorites.Favorites, err
	}

	defer helpers2.CloseBody(res.Body)

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
