package block_pods

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/common"
	"fmt"
	"net/http"
)

func UnlinkPodFromBlock(jwtToken string, podParam common.ResourceIdParam) error {
	route, err := helpers.GetRoute(golang.RouteBlockPodsUnlinkPodFromBlock, podParam.BlockId, podParam.PodId, podParam.KeyId)
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

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
