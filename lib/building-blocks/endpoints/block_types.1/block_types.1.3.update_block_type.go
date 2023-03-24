package block_types

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/request"
	"development/go/recipes/lib/building-blocks/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func UpdateBlockType(jwtToken string, reqBody request.BlockTypeReqBody, blockTypeId string) (response.BlockType, error) {
	resBlockType := response.BlockType{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resBlockType, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(building_blocks.RouteBlockTypesUpdateBlockType, blockTypeId)
	if err != nil {
		fmt.Println(err)
		return resBlockType, err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resBlockType, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlockType, err
	}

	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resBlockType, err
	}

	err = json.Unmarshal(body, &resBlockType)
	if err != nil {
		fmt.Println(err)
		return resBlockType, err
	}
	return resBlockType, nil
}