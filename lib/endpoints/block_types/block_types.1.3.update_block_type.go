package block_types

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func UpdateBlockType(jwtToken string, reqBody request.BlockTypeReqBody, blockTypeId string) (response.BlockType, error) {
	resBlockType := response.BlockType{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resBlockType, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(lib.RouteBlockTypesUpdateBlockType, blockTypeId)
	if err != nil {
		fmt.Println(err)
		return resBlockType, err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
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
