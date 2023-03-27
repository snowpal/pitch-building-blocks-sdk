package block_types

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/request"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func AddBlockType(jwtToken string, reqBody request.BlockTypeReqBody) (response.BlockType, error) {
	resBlockType := response.BlockType{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resBlockType, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(lib.RouteBlockTypesAddBlockType)
	if err != nil {
		fmt.Println(err)
		return resBlockType, err
	}

	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resBlockType, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlockType, err
	}

	defer helpers2.CloseBody(res.Body)

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
