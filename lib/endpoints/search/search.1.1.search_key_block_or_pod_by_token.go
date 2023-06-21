package search

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func SrchKeyBlockOrPodByToken(jwtToken string, searchToken string) ([]response.SearchResource, error) {
	resSearchResources := response.SearchResources{}
	route, err := helpers2.GetRoute(lib.RouteSearchSearchKeyBlockOrPodByToken, searchToken)
	if err != nil {
		return resSearchResources.Results, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resSearchResources.Results, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resSearchResources.Results, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resSearchResources.Results, err
	}

	err = json.Unmarshal(body, &resSearchResources)
	if err != nil {
		return resSearchResources.Results, err
	}
	return resSearchResources.Results, nil
}
