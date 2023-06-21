package projectKeys

import (
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
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
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, nil)
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
