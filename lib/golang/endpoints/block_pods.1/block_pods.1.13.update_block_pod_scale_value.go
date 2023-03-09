package block_pods

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/common"
	"development/go/recipes/lib/golang/structs/request"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func UpdateBlockPodScaleValue(jwtToken string, slimPod common.SlimPod, pod request.Pod) error {
	requestBody, err := helpers.GetRequestBody(pod)
	if err != nil {
		return err
	}
	payload := strings.NewReader(requestBody)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, helpers.GetRoute(golang.RouteBlockPodsUpdateBlockPodScaleValue,
		slimPod.ID, slimPod.Key.ID, slimPod.Block.ID), payload)
	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
