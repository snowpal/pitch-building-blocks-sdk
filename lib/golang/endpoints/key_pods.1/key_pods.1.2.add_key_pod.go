package key_pods_1

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

func AddKeyPod(jwtToken string, reqBody request.AddPodReqBody, keyId string) (response.Pod, error) {
	resPod := response.Pod{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}
	payload := strings.NewReader(requestBody)

	var route string
	route, err = helpers.GetRoute(golang.RouteKeyPodsAddKeyPod, keyId)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}

	defer helpers.CloseBody(res.Body)

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
