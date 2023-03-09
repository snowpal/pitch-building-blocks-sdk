package blocks_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type AddBlockReqBody struct {
	Name *string `json:"blockName"`
}

type BlockByTemplate struct {
	KeyId        string
	TemplateId   string
	ExcludeTasks *bool
}

func AddBlockBasedOnTemplate(jwtToken string, reqBody AddBlockReqBody, block BlockByTemplate) (response.Block, error) {
	response := response.Block{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return response, err
	}
	payload := strings.NewReader(requestBody)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, helpers.GetRoute(golang.RouteBlocksAddBlockBasedOnTemplate,
		block.KeyId, block.TemplateId, strconv.FormatBool(*block.ExcludeTasks)), payload)
	if err != nil {
		fmt.Println(err)
		return response, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return response, err
	}

	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return response, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
		return response, err
	}
	return response, nil
}
