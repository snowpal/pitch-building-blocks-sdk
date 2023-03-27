package keyPods

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/common"
	"development/go/recipes/lib/structs/request"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func UpdateKeyPod(
	jwtToken string,
	reqBody request.UpdatePodReqBody,
	podParam common.ResourceIdParam,
) (response.Pod, error) {
	resPod := response.Pod{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}
	payload := strings.NewReader(requestBody)

	var route string
	route, err = helpers2.GetRoute(lib.RouteKeyPodsUpdateKeyPod, podParam.PodId, podParam.KeyId)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}

	err = json.Unmarshal(body, &resPod)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}
	return resPod, nil
}
