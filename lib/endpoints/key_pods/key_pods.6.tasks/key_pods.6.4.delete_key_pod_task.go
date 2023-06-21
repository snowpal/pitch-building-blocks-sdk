package keyPods

import (
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
)

func DeleteKeyPodTask(jwtToken string, taskParam request.TaskIdParam) error {
	route, err := helpers.GetRoute(
		lib.RouteKeyPodsDeleteKeyPodTask,
		*taskParam.TaskId,
		taskParam.KeyId,
		*taskParam.PodId,
	)
	req, err := http.NewRequest(http.MethodDelete, route, nil)
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
