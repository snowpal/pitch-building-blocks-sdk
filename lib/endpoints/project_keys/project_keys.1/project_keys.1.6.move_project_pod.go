package projectKeys

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/request"
	"fmt"
	"net/http"
)

func MoveProjectPod(jwtToken string, projectPodParam request.CopyMoveProjectPodParam) error {
	route, err := helpers.GetRoute(
		lib.RouteProjectKeysMoveProjectPod,
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
	req, err = http.NewRequest(http.MethodPatch, route, nil)
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
