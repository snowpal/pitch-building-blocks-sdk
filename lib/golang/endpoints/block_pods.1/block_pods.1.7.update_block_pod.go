package block_pods

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

func RouteBlockPodsUpdateBlockPod(jwtToken string, slimPod common.SlimPod, pod request.Pod) (response.Pod, error) {
	podResponse := response.Pod{}

	requestBody, err := helpers.GetRequestBody(pod)
	if err != nil {
		fmt.Println(err)
		return podResponse, err
	}
	payload := strings.NewReader(requestBody)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, helpers.GetRoute(golang.RouteBlockPodsUpdateBlockPod, slimPod.ID, slimPod.Key.ID, slimPod.Block.ID), payload)
	if err != nil {
		fmt.Println(err)
		return podResponse, err
	}
	helpers.AddUserHeaders(jwtToken, req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return podResponse, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return podResponse, err
	}
	fmt.Println(string(body))

	err = json.Unmarshal(body, &podResponse)
	if err != nil {
		fmt.Println(err)
		return podResponse, err
	}
	return podResponse, nil
}
