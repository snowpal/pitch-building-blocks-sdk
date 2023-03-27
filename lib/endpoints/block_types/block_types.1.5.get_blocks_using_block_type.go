package block_types

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetBlocksUsingBlockType(jwtToken string, blockTypeId string) ([]response.Block, error) {
	resBlocks := response.Blocks{}
	route, err := helpers2.GetRoute(lib.RouteBlockTypesGetBlocksUsingBlockType, blockTypeId)
	if err != nil {
		fmt.Println(err)
		return resBlocks.Blocks, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resBlocks.Blocks, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlocks.Blocks, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
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
