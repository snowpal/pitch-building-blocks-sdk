package blocks_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/common"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetBlocksLinkedToPods(jwtToken string, blockParam common.ResourceIdParam) ([]response.Block, error) {
	resBlocks := response.Blocks{}
	route, err := helpers.GetRoute(
		golang.RouteBlocksGetBlocksLinkedToPod,
		blockParam.PodId,
		blockParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resBlocks.Blocks, err
	}

	req, _ := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resBlocks.Blocks, err
	}

	helpers.AddUserHeaders(jwtToken, req)
	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlocks.Blocks, err
	}

	defer helpers.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resBlocks.Blocks, err
	}

	err = json.Unmarshal(body, &resBlocks)
	if err != nil {
		fmt.Println(err)
		return resBlocks.Blocks, err
	}
	return resBlocks.Blocks, nil
}
