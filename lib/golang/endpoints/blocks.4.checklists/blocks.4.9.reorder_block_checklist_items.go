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

func ReorderBlockChecklistItems(
	jwtToken string,
	reqBody request.ReorderChecklistItemsReqBody,
	checklistParam request.ChecklistIdParam,
) ([]response.ChecklistItem, error) {
	resChecklistItems := response.ChecklistItems{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resChecklistItems.ChecklistItems, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(
		golang.RouteBlocksReorderBlockChecklistItems,
		*checklistParam.BlockId,
		*checklistParam.ChecklistId,
		checklistParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resChecklistItems.ChecklistItems, err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resChecklistItems.ChecklistItems, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resChecklistItems.ChecklistItems, err
	}

	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resChecklistItems.ChecklistItems, err
	}

	err = json.Unmarshal(body, &resChecklistItems)
	if err != nil {
		fmt.Println(err)
		return resChecklistItems.ChecklistItems, err
	}
	return resChecklistItems.ChecklistItems, nil
}
