package blocks_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetBlocksAvailableToBeLinkedToThisKey(jwtToken string, keyId string) ([]response.Block, error) {
	resBlocks := response.Blocks{}
	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteBlocksGetBlocksAvailableToBeLinkedToThisKey, keyId)
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
	res, _ := client.Do(req)
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
