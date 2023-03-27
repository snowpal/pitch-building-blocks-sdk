package scales

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

type GetScalesParam struct {
	IncludeCounts bool
	ExcludeEmpty  bool
}

func GetScales(jwtToken string, scalesParam GetScalesParam) ([]response.Scale, error) {
	resScales := response.Scales{}
	route, err := helpers2.GetRoute(
		lib.RouteScalesGetScales,
		strconv.FormatBool(scalesParam.IncludeCounts),
		strconv.FormatBool(scalesParam.ExcludeEmpty),
	)
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resScales.Scales, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resScales.Scales, err
	}

	defer helpers2.CloseBody(res.Body)

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
