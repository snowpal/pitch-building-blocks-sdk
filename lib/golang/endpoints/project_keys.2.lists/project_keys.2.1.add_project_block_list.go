package project_keys_2

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

func AddProjectBlockList(
	jwtToken string,
	reqBody request.AddProjectListReqBody,
	projectListParam common.ResourceIdParam,
) (response.ProjectList, error) {
	resProjectList := response.ProjectList{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resProjectList, err
	}
	payload := strings.NewReader(requestBody)
	client := &http.Client{}

	var route string
	route, err = helpers.GetRoute(
		golang.RouteProjectKeysAddProjectBlockList,
		projectListParam.BlockId,
		projectListParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resProjectList, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resProjectList, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resProjectList, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resProjectList, err
	}

	err = json.Unmarshal(body, &resProjectList)
	if err != nil {
		fmt.Println(err)
		return resProjectList, err
	}
	return resProjectList, nil
}
