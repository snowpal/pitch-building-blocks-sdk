package scales

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func AddScale(jwtToken string, reqBody request.ScaleReqBody) (response.Scale, error) {
	resScale := response.Scale{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		return resScale, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(lib.RouteScalesAddScale)
	if err != nil {
		return resScale, err
	}

	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		return resScale, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resScale, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resScale, err
	}

	err = json.Unmarshal(body, &resScale)
	if err != nil {
		return resScale, err
	}
	return resScale, nil
}
