package project_keys_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"fmt"
	"net/http"
)

func LinkProjectPodToBlock(jwtToken string, podParam request.ProjectListIdParam) error {
	client := &http.Client{}
	route, err := helpers.GetRoute(
		golang.RouteProjectKeysLinkProjectPodToBlock,
		podParam.BlockId,
		*podParam.PodId,
		podParam.KeyId,
		podParam.ProjectListId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
