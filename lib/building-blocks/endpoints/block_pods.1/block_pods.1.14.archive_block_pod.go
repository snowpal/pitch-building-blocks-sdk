package block_pods

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/common"
	"fmt"
	"net/http"
)

func ArchiveBlockPod(jwtToken string, podParam common.ResourceIdParam) error {
	route, err := helpers.GetRoute(
		building_blocks.RouteBlockPodsArchiveBlockPod,
		podParam.PodId,
		podParam.KeyId,
		podParam.BlockId,
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
