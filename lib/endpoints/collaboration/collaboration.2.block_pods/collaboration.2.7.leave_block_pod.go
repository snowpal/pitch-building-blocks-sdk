package collaboration

import (
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
)

func LeaveBlockPod(jwtToken string, podParam common.ResourceIdParam) error {
	route, err := helpers.GetRoute(
		lib.RouteCollaborationLeaveBlockPod,
		podParam.PodId,
		podParam.KeyId,
		podParam.BlockId,
	)
	req, err := http.NewRequest(http.MethodPatch, route, nil)
	if err != nil {
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		return err
	}
	return nil
}
