package templates

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetBlockTemplates(jwtToken string) ([]response.BlockTemplate, error) {
	resBlockTemplates := response.BlockTemplates{}
	route, err := helpers2.GetRoute(lib.RouteTemplatesGetBlockTemplates)
	if err != nil {
		return resBlockTemplates.Templates, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resBlockTemplates.Templates, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resBlockTemplates.Templates, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resBlockTemplates.Templates, err
	}

	err = json.Unmarshal(body, &resBlockTemplates)
	if err != nil {
		return resBlockTemplates.Templates, err
	}
	return resBlockTemplates.Templates, nil
}
