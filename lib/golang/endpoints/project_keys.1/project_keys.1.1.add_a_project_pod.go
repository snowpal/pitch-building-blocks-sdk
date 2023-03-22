package project_keys_1

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

func AddProjectPod(
	jwtToken string,
	reqBody request.AddPodReqBody,
	projectListParam request.ProjectListIdParam,
) (response.Pod, error) {
	resProjectPod := response.Pod{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resProjectPod, err
	}
	payload := strings.NewReader(requestBody)

	var route string
	route, err = helpers.GetRoute(
		golang.RouteProjectKeysAddAProjectPod,
		projectListParam.BlockId,
		projectListParam.KeyId,
		projectListParam.ProjectListId,
	)
	if err != nil {
		fmt.Println(err)
		return resProjectPod, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resProjectPod, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resProjectPod, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resProjectPod, err
	}

	err = json.Unmarshal(body, &resProjectPod)
	if err != nil {
		fmt.Println(err)
		return resProjectPod, err
	}
	return resProjectPod, nil
}
