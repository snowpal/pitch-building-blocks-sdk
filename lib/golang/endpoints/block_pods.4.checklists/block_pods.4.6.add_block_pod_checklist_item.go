package block_pods_4

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

func AddBlockPodChecklistItem(
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
		golang.RouteBlockPodsAddBlockPodChecklistItem,
		*checklistParam.PodId,
		*checklistParam.ChecklistId,
		checklistParam.KeyId,
		*checklistParam.BlockId,
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
