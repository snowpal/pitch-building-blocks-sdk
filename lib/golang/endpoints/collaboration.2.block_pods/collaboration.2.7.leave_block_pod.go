package collaboration_2

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/common"
	"fmt"
	"net/http"
)

func LeaveBlockPod(jwtToken string, podParam common.ResourceIdParam) error {
	client := &http.Client{}
	route, err := helpers.GetRoute(
		golang.RouteCollaborationLeaveBlockPod,
		podParam.PodId,
		podParam.KeyId,
		podParam.BlockId,
	)
	req, err := http.NewRequest(http.MethodPatch, route, nil)
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
