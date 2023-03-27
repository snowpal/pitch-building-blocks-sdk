package version

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetLatestVersion() (response.Version, error) {
	resVersion := response.Version{}
	route, err := helpers2.GetRoute(lib.RouteVersionGetLatestVersion)
	if err != nil {
		fmt.Println(err)
		return resVersion, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resVersion, err
	}

	helpers2.AddAppHeaders(req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resVersion, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resVersion, err
	}

	err = json.Unmarshal(body, &resVersion)
	if err != nil {
		fmt.Println(err)
		return resVersion, err
	}
	return resVersion, nil
}
