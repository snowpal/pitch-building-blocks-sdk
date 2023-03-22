package block_types

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

func AddBlockType(jwtToken string, reqBody request.BlockTypeReqBody) (response.BlockType, error) {
	resBlockType := response.BlockType{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resBlockType, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(golang.RouteBlockTypesAddBlockType)
	if err != nil {
		fmt.Println(err)
		return resBlockType, err
	}

	req, err := http.NewRequest(http.MethodPost, route, payload)
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
