package keyPods

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/common"
	"fmt"
	"net/http"
)

func LinkPodToKey(jwtToken string, podParam common.ResourceIdParam) error {
	route, err := helpers.GetRoute(lib.RouteKeyPodsLinkKeyPodToKey, podParam.KeyId, podParam.PodId)
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
