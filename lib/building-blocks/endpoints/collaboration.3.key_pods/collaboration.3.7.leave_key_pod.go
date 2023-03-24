package collaboration_3

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/common"
	"fmt"
	"net/http"
)

func LeaveKeyPod(jwtToken string, podParam common.ResourceIdParam) error {
	route, err := helpers.GetRoute(
		building_blocks.RouteCollaborationLeaveKeyPod,
		podParam.PodId,
		podParam.KeyId,
	)
	req, err := http.NewRequest(http.MethodPatch, route, nil)
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
