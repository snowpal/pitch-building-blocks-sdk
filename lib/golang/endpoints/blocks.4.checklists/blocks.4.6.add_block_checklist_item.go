package blocks_4

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func AddBlockChecklistItem(
	jwtToken string,
	reqBody request.ChecklistItemReqBody,
	checklistParam request.ChecklistIdParam,
) (response.ChecklistItem, error) {
	resChecklistItem := response.ChecklistItem{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resChecklistItem, err
	}
	payload := strings.NewReader(requestBody)
	client := &http.Client{}
	route, err := helpers.GetRoute(
		golang.RouteBlocksAddBlockChecklistItem,
		*checklistParam.BlockId,
		*checklistParam.ChecklistId,
		checklistParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resChecklistItem, err
	}

	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resChecklistItem, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resChecklistItem, err
	}
	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resChecklistItem, err
	}

	err = json.Unmarshal(body, &resChecklistItem)
	if err != nil {
		fmt.Println(err)
		return resChecklistItem, err
	}
	return resChecklistItem, nil
}