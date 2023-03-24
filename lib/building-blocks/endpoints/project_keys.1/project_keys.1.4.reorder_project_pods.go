package project_keys_1

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/common"
	"development/go/recipes/lib/building-blocks/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ReorderProjectPodsReqBody struct {
	ProjectListId       string `json:"sourceProjectListId"`
	ProjectPodIds       string `json:"sourceProjectListPodIds"`
	TargetProjectListId string `json:"targetProjectListId"`
	TargetProjectPodIds string `json:"targetProjectListPodIds"`
}

func ReorderProjectPods(
	jwtToken string,
	reqBody ReorderProjectPodsReqBody,
	podParam common.ResourceIdParam,
) ([]response.Pod, error) {
	resProjectPods := response.Pods{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resProjectPods.Pods, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(
		building_blocks.RouteProjectKeysReorderProjectPods,
		podParam.BlockId,
		podParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resProjectPods.Pods, err
	}
	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resProjectPods.Pods, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resProjectPods.Pods, err
	}

	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resProjectPods.Pods, err
	}

	err = json.Unmarshal(body, &resProjectPods)
	if err != nil {
		fmt.Println(err)
		return resProjectPods.Pods, err
	}
	return resProjectPods.Pods, nil
}
