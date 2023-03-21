package scales

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetBlocksUsingScale(jwtToken string, scaleId string) ([]response.Block, error) {
	resBlocks := response.Blocks{}
	route, err := helpers.GetRoute(golang.RouteScalesGetBlocksUsingScale, scaleId)
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

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlocks.Blocks, err
	}

	defer helpers.CloseBody(res.Body)

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
