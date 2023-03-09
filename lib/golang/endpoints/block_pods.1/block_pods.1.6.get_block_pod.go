package block_pods

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/common"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetBlockPod(jwtToken string, pod common.SlimPod) (response.Pod, error) {
	podResponse := response.Pod{}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, helpers.GetRoute(golang.RouteBlockPodsGetBlockPod, pod.ID, pod.Key.ID, pod.Block.ID), nil)
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

	err = json.Unmarshal(body, &podResponse)
	if err != nil {
		fmt.Println(err)
		return podResponse, err
	}
	return podResponse, nil
}
