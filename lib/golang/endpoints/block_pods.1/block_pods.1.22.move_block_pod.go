package block_pods

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"fmt"
	"net/http"
)

func MoveBlockPod(jwtToken string, podParam request.CopyMovePodParam) error {
	route, err := helpers.GetRoute(
		golang.RouteBlockPodsMoveBlockPod,
		podParam.PodId,
		podParam.KeyId,
		podParam.BlockId,
		podParam.TargetKeyId,
		podParam.TargetBlockId,
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
