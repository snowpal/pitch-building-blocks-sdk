package blocks

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

type UpdateBlockReqBody struct {
	Name              *string `json:"blockName"`
	BlockId           *string `json:"blockId"`
	SimpleDescription *string `json:"simpleDescription"`
	DueDate           *string `json:"blockDueDate"`
	StartTime         *string `json:"blockStartTime"`
	EndTime           *string `json:"blockEndTime"`
	Color             *string `json:"blockColor"`
	Tags              *string `json:"blockTags"`
	KanbanMode        *bool   `json:"kanbanMode"`
	Completed         bool    `json:"blockCompleted"`
}

func UpdateBlock(jwtToken string, reqBody UpdateBlockReqBody, blockParam common.ResourceIdParam) (response.Block, error) {
	resBlock := response.Block{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(lib.RouteBlocksUpdateBlock, blockParam.BlockId, blockParam.KeyId)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
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
