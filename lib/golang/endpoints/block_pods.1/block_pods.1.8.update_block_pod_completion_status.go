package block_pods

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/common"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func UpdateBlockPodCompletionStatus(jwtToken string, slimPod common.SlimPod, pod request.Pod) (response.Pod, error) {
	podResponse := response.Pod{}

	requestBody, err := helpers.GetRequestBody(pod)
	if err != nil {
		fmt.Println(err)
		return podResponse, err
	}
	payload := strings.NewReader(requestBody)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, helpers.GetRoute(golang.RouteBlockPodsUpdateBlockPodCompletionStatus, slimPod.ID, slimPod.Key.ID, slimPod.Block.ID), payload)
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
	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return podResponse, err
	}
	return podResponse, nil
}
