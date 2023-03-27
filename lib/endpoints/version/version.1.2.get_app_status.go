package version

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetAppStatus() (response.Version, error) {
	resAppStatus := response.Version{}
	route, err := helpers2.GetRoute(lib.RouteVersionGetAppStatus)
	if err != nil {
		fmt.Println(err)
		return resAppStatus, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resAppStatus, err
	}

	helpers2.AddAppHeaders(req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resAppStatus, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resAppStatus, err
	}

	err = json.Unmarshal(body, &resAppStatus)
	if err != nil {
		fmt.Println(err)
		return resAppStatus, err
	}
	return resAppStatus, nil
}
