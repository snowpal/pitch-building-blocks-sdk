package blocks_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"fmt"
	"net/http"
)

type AddBlockTypeIdParam struct {
	KeyId       string
	BlockId     string
	BlockTypeId string
}

func AddPodTypeToBlockPod(jwtToken string, podParam AddBlockTypeIdParam) error {
	route, err := helpers.GetRoute(
		golang.RouteBlocksAddBlockTypeToBlock,
		podParam.BlockId,
		podParam.BlockTypeId,
		podParam.KeyId,
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
