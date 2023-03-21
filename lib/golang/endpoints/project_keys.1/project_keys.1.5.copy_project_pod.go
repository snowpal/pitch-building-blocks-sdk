package project_keys_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"fmt"
	"net/http"
)

func CopyProjectPod(jwtToken string, projectPodParam request.CopyMoveProjectPodParam) error {
	route, err := helpers.GetRoute(
		golang.RouteProjectKeysCopyProjectPod,
		projectPodParam.PodId,
		projectPodParam.KeyId,
		projectPodParam.BlockId,
		projectPodParam.TargetKeyId,
		projectPodParam.TargetBlockId,
		projectPodParam.TargetProjectListId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
