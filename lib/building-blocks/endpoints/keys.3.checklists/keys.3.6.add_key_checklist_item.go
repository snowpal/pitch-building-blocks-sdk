package keys_3

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/request"
	"development/go/recipes/lib/building-blocks/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func AddKeyChecklistItem(
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
	route, err := helpers.GetRoute(
		building_blocks.RouteKeysAddKeyChecklistItem,
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

	res, err := helpers.MakeRequest(req)
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