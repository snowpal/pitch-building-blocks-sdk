package blocks_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/common"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func UpdateBlockScaleValue(
	jwtToken string,
	reqBody request.UpdateScaleValueReqBody,
	blockParam common.ResourceIdParam,
) (response.UpdateBlockScaleValue, error) {
	resBlockScaleValue := response.UpdateBlockScaleValue{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resBlockScaleValue, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(
		golang.RouteBlocksUpdateBlockScaleValue,
		blockParam.BlockId,
		blockParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resBlockScaleValue, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resBlockScaleValue, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlockScaleValue, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resBlockScaleValue, err
	}

	err = json.Unmarshal(body, &resBlockScaleValue)
	if err != nil {
		fmt.Println(err)
		return resBlockScaleValue, err
	}
	return resBlockScaleValue, nil
}
