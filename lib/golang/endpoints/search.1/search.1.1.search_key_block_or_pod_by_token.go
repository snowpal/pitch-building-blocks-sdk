package search

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SrchKeyBlockOrPodByToken(jwtToken string, searchToken string) ([]response.SearchResource, error) {
	resSearchResources := response.SearchResources{}
	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteSearchSearchKeyBlockOrPodByToken, searchToken)
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

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resSearchResources.Results, err
	}

	defer helpers.CloseBody(res.Body)

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
