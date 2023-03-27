package projectKeys

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	request2 "development/go/recipes/lib/structs/request"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func AddProjectPod(
	jwtToken string,
	reqBody request2.AddPodReqBody,
	projectListParam request2.ProjectListIdParam,
) (response.Pod, error) {
	resProjectPod := response.Pod{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resProjectPod, err
	}
	payload := strings.NewReader(requestBody)

	var route string
	route, err = helpers2.GetRoute(
		lib.RouteProjectKeysAddAProjectPod,
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

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resProjectPod, err
	}

	defer helpers2.CloseBody(res.Body)

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
