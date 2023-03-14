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
	"strings"
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
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}
	payload := strings.NewReader(requestBody)
	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteBlocksUpdateBlock, blockParam.BlockId, blockParam.KeyId)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resBlock, err
	}
	defer helpers.CloseBody(res.Body)

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
