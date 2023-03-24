package keys_1

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

type KeyByTemplateParam struct {
	TemplateId    string
	ExcludeBlocks bool
	ExcludePods   bool
	ExcludeTasks  bool
}

func AddKeyBasedOnTemplate(
	jwtToken string,
	reqBody request.AddKeyReqBody,
	keyParam KeyByTemplateParam,
) (response.Key, error) {
	resKey := response.Key{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resKey, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(
		building_blocks.RouteKeysAddKeyBasedOnTemplate,
		keyParam.TemplateId,
		strconv.FormatBool(keyParam.ExcludeBlocks),
		strconv.FormatBool(keyParam.ExcludePods),
		strconv.FormatBool(keyParam.ExcludeTasks),
	)
	if err != nil {
		fmt.Println(err)
		return resKey, err
	}
	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resKey, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resKey, err
	}

	defer helpers.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resKey, err
	}

	err = json.Unmarshal(body, &resKey)
	if err != nil {
		fmt.Println(err)
		return resKey, err
	}
	return resKey, nil
}
