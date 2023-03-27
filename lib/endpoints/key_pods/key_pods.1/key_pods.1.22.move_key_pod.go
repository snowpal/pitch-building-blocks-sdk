package keyPods

import (
	"fmt"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
)

func MoveKeyPod(jwtToken string, podParam request.CopyMovePodParam) error {
	route, err := helpers.GetRoute(
		lib.RouteKeyPodsMoveKeyPod,
		podParam.PodId,
		podParam.KeyId,
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
