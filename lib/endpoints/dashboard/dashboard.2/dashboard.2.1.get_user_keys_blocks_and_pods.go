package dashboard

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetUserKeysBlocksAndPods(jwtToken string) ([]response.UserKey, error) {
	resUserKeys := response.UserKeys{}
	route, err := helpers2.GetRoute(lib.RouteDashboardGetUserKeysBlocksAndPods)
	if err != nil {
		return resUserKeys.Keys, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resUserKeys.Keys, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resUserKeys.Keys, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resUserKeys.Keys, err
	}

	err = json.Unmarshal(body, &resUserKeys)
	if err != nil {
		return resUserKeys.Keys, err
	}
	return resUserKeys.Keys, nil
}
