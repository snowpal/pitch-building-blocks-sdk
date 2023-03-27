package scales

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetScale(jwtToken string, scaleId string) (response.Scale, error) {
	resScale := response.Scale{}
	route, err := helpers2.GetRoute(lib.RouteScalesGetScale, scaleId)
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resScale, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resScale, err
	}

	defer helpers2.CloseBody(res.Body)

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
