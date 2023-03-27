package search

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SrchKeyBlockOrPodByToken(jwtToken string, searchToken string) ([]response.SearchResource, error) {
	resSearchResources := response.SearchResources{}
	route, err := helpers2.GetRoute(lib.RouteSearchSearchKeyBlockOrPodByToken, searchToken)
	if err != nil {
		fmt.Println(err)
		return resSearchResources.Results, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resSearchResources.Results, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resSearchResources.Results, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resSearchResources.Results, err
	}

	err = json.Unmarshal(body, &resSearchResources)
	if err != nil {
		fmt.Println(err)
		return resSearchResources.Results, err
	}
	return resSearchResources.Results, nil
}
