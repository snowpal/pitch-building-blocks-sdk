package favorites

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func AddKeyAsFavorite(jwtToken string, keyId string) (response.AddFavorite, error) {
	resFavorite := response.AddFavorite{}
	route, err := helpers2.GetRoute(lib.RouteFavoritesAddKeyAsFavorite, keyId)
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

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resFavorite, err
	}

	defer helpers2.CloseBody(res.Body)

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
