package blocks_1

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/request"
	"development/go/recipes/lib/building-blocks/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type BlockByTemplateParam struct {
	KeyId        string
	TemplateId   string
	ExcludePods  bool
	ExcludeTasks bool
}

func AddBlockBasedOnTemplate(
	jwtToken string,
	reqBody request.AddBlockReqBody,
	blockParam BlockByTemplateParam,
) (response.Block, error) {
	resBlock := response.Block{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}
	payload := strings.NewReader(requestBody)
	var route string
	route, err = helpers.GetRoute(
		building_blocks.RouteBlocksAddBlockBasedOnTemplate,
		blockParam.KeyId,
		blockParam.TemplateId,
		strconv.FormatBool(blockParam.ExcludePods),
		strconv.FormatBool(blockParam.ExcludeTasks),
	)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}
	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	err = json.Unmarshal(body, &resBlock)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}
	return resBlock, nil
}