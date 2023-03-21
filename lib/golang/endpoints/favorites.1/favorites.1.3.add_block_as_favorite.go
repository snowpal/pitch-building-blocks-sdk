package favorites

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/common"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func AddBlockAsFavorite(jwtToken string, favoriteParam common.ResourceIdParam) (response.AddFavorite, error) {
	resFavorite := response.AddFavorite{}
	route, err := helpers.GetRoute(
		golang.RouteFavoritesAddBlockAsFavorite,
		favoriteParam.BlockId,
		favoriteParam.KeyId,
	)
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
