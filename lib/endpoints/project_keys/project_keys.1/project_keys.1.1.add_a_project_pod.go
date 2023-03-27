package projectKeys

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
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
