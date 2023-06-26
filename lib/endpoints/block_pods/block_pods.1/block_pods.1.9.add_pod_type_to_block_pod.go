package blockPods

import (
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
)

func AddPodTypeToBlockPod(jwtToken string, podParam request.AddPodTypeIdParam) error {
	route, err := helpers.GetRoute(
		lib.RouteBlockPodsAddPodTypeToBlockPod,
		podParam.PodId,
		podParam.PodTypeId,
		podParam.KeyId,
		*podParam.BlockId,
	)
	if err != nil {
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, nil)
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
