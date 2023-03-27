package blocks

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/common"
	"fmt"
	"net/http"
)

func UnarchiveBlock(jwtToken string, blockParam common.ResourceIdParam) error {
	route, err := helpers.GetRoute(
		lib.RouteBlocksUnarchiveBlock,
		blockParam.BlockId,
		blockParam.KeyId,
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
