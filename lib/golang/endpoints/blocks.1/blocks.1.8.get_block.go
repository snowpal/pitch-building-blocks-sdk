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

func GetBlock(jwtToken string, block common.SlimBlock) (response.Block, error) {
	blockResponse := response.Block{}
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, helpers.GetRoute(golang.RouteBlocksGetBlock, block.ID, block.Key.ID),
		nil)
	if err != nil {
		fmt.Println(err)
		return blockResponse, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return blockResponse, err
	}

	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return blockResponse, err
	}

	err = json.Unmarshal(body, &blockResponse)
	if err != nil {
		fmt.Println(err)
		return blockResponse, err
	}
	return blockResponse, nil
}
