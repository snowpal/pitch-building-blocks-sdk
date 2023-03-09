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

func LinkPodToBlock(jwtToken string, block common.SlimBlock, pod request.Pod) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, helpers.GetRoute(golang.RouteBlockPodsLinkPodToBlock, block.ID, *pod.ID, block.Key.ID), nil)

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
