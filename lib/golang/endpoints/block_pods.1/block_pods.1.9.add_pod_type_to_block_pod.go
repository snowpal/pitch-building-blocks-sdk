package block_pods

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/common"
	"development/go/recipes/lib/golang/structs/request"
	"fmt"
	"io"
	"net/http"
)

func AddPodTypeToBlockPod(jwtToken string, slimPod common.SlimPod, podType request.PodType) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, helpers.GetRoute(golang.RouteBlockPodsAddPodTypeToBlockPod, slimPod.ID, *podType.ID, slimPod.Key.ID, slimPod.Block.ID), nil)
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
