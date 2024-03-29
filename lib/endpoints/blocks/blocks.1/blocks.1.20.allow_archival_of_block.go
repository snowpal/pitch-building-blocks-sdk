package blocks

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func AllowArchivalOfBlock(
	jwtToken string,
	reqBody common.AllowArchivalReqBody,
	blockParam common.ResourceIdParam,
) (response.Block, error) {
	resBlock := response.Block{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		return resBlock, err
	}
	payload := strings.NewReader(requestBody)

	var route string
	route, err = helpers2.GetRoute(
		lib.RouteBlocksAllowArchivalOfBlock,
		blockParam.BlockId,
		blockParam.KeyId,
	)
	if err != nil {
		return resBlock, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		return resBlock, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resBlock, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resBlock, err
	}

	err = json.Unmarshal(body, &resBlock)
	if err != nil {
		return resBlock, err
	}
	return resBlock, nil
}
