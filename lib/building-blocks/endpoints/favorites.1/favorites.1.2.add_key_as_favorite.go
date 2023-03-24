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

func AddKeyAsFavorite(jwtToken string, keyId string) (response.AddFavorite, error) {
	resFavorite := response.AddFavorite{}
	route, err := helpers.GetRoute(building_blocks.RouteFavoritesAddKeyAsFavorite, keyId)
	if err != nil {
		fmt.Println(err)
		return resFavorite, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, nil)
	if err != nil {
		fmt.Println(err)
		return resFavorite, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resFavorite, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resFavorite, err
	}

	err = json.Unmarshal(body, &resFavorite)
	if err != nil {
		fmt.Println(err)
		return resFavorite, err
	}
	return resFavorite, nil
}
