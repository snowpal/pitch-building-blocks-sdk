package keys_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"
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
	client := &http.Client{}
	route, err := helpers.GetRoute(
		golang.RouteKeysAddKeyBasedOnTemplate,
		keyParam.TemplateId,
		strconv.FormatBool(keyParam.ExcludeBlocks),
		strconv.FormatBool(keyParam.ExcludePods),
		strconv.FormatBool(keyParam.ExcludeTasks),
	)
	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resKey, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, _ := client.Do(req)
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
