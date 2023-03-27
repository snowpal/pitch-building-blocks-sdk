package projectKeys

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/request"
	"fmt"
	"net/http"
)

func CopyProjectPod(jwtToken string, projectPodParam request.CopyMoveProjectPodParam) error {
	route, err := helpers.GetRoute(
		lib.RouteProjectKeysCopyProjectPod,
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

	_, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
