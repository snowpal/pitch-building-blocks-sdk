package dashboard

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetBlocksAndPodsBasedOnScales(jwtToken string) ([]response.ScalesKey, error) {
	resScalesKeys := response.ScalesKeys{}
	route, err := helpers2.GetRoute(lib.RouteDashboardGetBlocksAndPodsBasedOnScales)
	if err != nil {
		return *resScalesKeys.Keys, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return *resScalesKeys.Keys, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return *resScalesKeys.Keys, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return *resScalesKeys.Keys, err
	}

	err = json.Unmarshal(body, &resScalesKeys)
	if err != nil {
		return *resScalesKeys.Keys, err
	}
	return *resScalesKeys.Keys, nil
}
