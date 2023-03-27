package projectKeys

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
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
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resProjectPods.Pods, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteProjectKeysReorderProjectPods,
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

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resProjectPods.Pods, err
	}

	defer helpers2.CloseBody(res.Body)

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
