package blocks_1

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"fmt"
	"net/http"
	"strings"
)

type BulkArchiveBlocksReqBody struct {
	BlockIds string `json:"blockIds"`
}

func BulkArchiveBlocks(jwtToken string, reqBody BulkArchiveBlocksReqBody, keyId string) error {
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return err
	}
	payload := strings.NewReader(requestBody)

	var route string
	route, err = helpers.GetRoute(building_blocks.RouteBlocksBulkArchiveBlocks, keyId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
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
