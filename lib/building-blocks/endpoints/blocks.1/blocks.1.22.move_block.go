package blocks_1

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/request"
	"fmt"
	"net/http"
)

func MoveBlock(jwtToken string, blockParam request.CopyMoveBlockParam) error {
	route, err := helpers.GetRoute(
		building_blocks.RouteBlocksMoveBlock,
		blockParam.BlockId,
		blockParam.KeyId,
		blockParam.TargetKeyId,
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
