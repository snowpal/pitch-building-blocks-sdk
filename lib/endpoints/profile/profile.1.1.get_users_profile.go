package profiles

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetUsersProfile(jwtToken string) (response.Profile, error) {
	resProfile := response.Profile{}
	route, err := helpers.GetRoute(lib.RouteProfileGetUsersProfile)
	if err != nil {
		return resProfile, err
	}
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resProfile, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		return resProfile, err
	}

	defer helpers.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		return resProfile, err
	}

	err = json.Unmarshal(body, &resProfile)
	if err != nil {
		return resProfile, err
	}
	return resProfile, nil
}
