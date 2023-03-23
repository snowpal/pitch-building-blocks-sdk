package scales

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type GetScalesParam struct {
	IncludeCounts bool
	ExcludeEmpty  bool
}

func GetScales(jwtToken string, scalesParam GetScalesParam) ([]response.Scale, error) {
	resScales := response.Scales{}
	route, err := helpers.GetRoute(
		golang.RouteScalesGetScales,
		strconv.FormatBool(scalesParam.IncludeCounts),
		strconv.FormatBool(scalesParam.ExcludeEmpty),
	)
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resScales.Scales, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resScales.Scales, err
	}

	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resScales.Scales, err
	}

	err = json.Unmarshal(body, &resScales)
	if err != nil {
		fmt.Println(err)
		return resScales.Scales, err
	}
	return resScales.Scales, nil
}
