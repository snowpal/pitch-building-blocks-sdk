package blockPods

import (
	"net/http"
	"strconv"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
)

func CopyBlockPod(jwtToken string, podParam request.CopyMovePodParam) error {
	route, err := helpers.GetRoute(
		lib.RouteBlockPodsCopyBlockPod,
		podParam.PodId,
		podParam.KeyId,
		podParam.BlockId,
		strconv.FormatBool(podParam.AllTasks),
		strconv.FormatBool(podParam.AllChecklists),
		podParam.TargetKeyId,
		podParam.TargetBlockId,
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
