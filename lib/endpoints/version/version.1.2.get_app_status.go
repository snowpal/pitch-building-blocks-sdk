package version

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetAppStatus() (response.Version, error) {
	resAppStatus := response.Version{}
	route, err := helpers2.GetRoute(lib.RouteVersionGetAppStatus)
	if err != nil {
		return resAppStatus, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resAppStatus, err
	}

	helpers2.AddAppHeaders(req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resAppStatus, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resAppStatus, err
	}

	err = json.Unmarshal(body, &resAppStatus)
	if err != nil {
		return resAppStatus, err
	}
	return resAppStatus, nil
}
