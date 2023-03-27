package blocks

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetBlocks(jwtToken string, blockParam request.GetBlocksParam) ([]response.Block, error) {
	resBlocks := response.Blocks{}
	route, err := helpers2.GetRoute(
		lib.RouteBlocksGetBlocks,
		blockParam.KeyId,
		blockParam.Filter,
		strconv.Itoa(blockParam.BatchIndex),
		strconv.FormatBool(blockParam.WriteOrHigherAcl),
	)
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
