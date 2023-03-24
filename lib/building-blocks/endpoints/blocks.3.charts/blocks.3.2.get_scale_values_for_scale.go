package blocks_3

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/request"
	"development/go/recipes/lib/building-blocks/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetScaleValuesForScale(jwtToken string, scaleParam request.ScaleIdParam) (response.ScaleValues, error) {
	resScaleValues := response.ScaleValues{}
	route, err := helpers.GetRoute(
		building_blocks.RouteBlocksGetScaleValuesForScale,
		scaleParam.KeyId,
		*scaleParam.BlockId,
		scaleParam.ScaleId,
	)
	if err != nil {
		fmt.Println(err)
		return resScaleValues, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resScaleValues, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resScaleValues, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resScaleValues, err
	}

	err = json.Unmarshal(body, &resScaleValues)
	if err != nil {
		fmt.Println(err)
		return resScaleValues, err
	}
	return resScaleValues, nil
}
