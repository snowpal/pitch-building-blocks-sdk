package blocks

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func CopyBlock(jwtToken string, blockParam request.CopyMoveBlockParam) (response.Block, error) {
	var resBlock response.Block
	route, err := helpers2.GetRoute(
		lib.RouteBlocksCopyBlock,
		blockParam.BlockId,
		blockParam.KeyId,
		strconv.FormatBool(blockParam.AllTasks),
		strings.Join(blockParam.PodIds, ","),
		strconv.FormatBool(blockParam.AllPods),
		strconv.FormatBool(blockParam.AllChecklists),
		blockParam.TargetKeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, nil)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	err = json.Unmarshal(body, &resBlock)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}
	return resBlock, nil
}
