package scales

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func AddScale(jwtToken string, reqBody request.ScaleReqBody) (response.Scale, error) {
	resScale := response.Scale{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resScale, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(golang.RouteScalesAddScale)
	if err != nil {
		fmt.Println(err)
		return resScale, err
	}

	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resScale, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resScale, err
	}

	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resScale, err
	}

	err = json.Unmarshal(body, &resScale)
	if err != nil {
		fmt.Println(err)
		return resScale, err
	}
	return resScale, nil
}
