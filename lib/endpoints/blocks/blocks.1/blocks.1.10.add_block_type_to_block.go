package blocks

import (
	"fmt"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
)

type AddBlockTypeIdParam struct {
	KeyId       string
	BlockId     string
	BlockTypeId string
}

func AddPodTypeToBlockPod(jwtToken string, podParam AddBlockTypeIdParam) error {
	route, err := helpers.GetRoute(
		lib.RouteBlocksAddBlockTypeToBlock,
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
