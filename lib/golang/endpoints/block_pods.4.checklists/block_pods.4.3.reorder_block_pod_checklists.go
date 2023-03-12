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

func ReorderBlockPodChecklists(
	jwtToken string,
	reqBody request.ReorderChecklistsReqBody,
	checklistParam request.ChecklistIdParam,
) ([]response.Checklist, error) {
	resChecklists := response.Checklists{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resChecklists.Checklists, err
	}
	payload := strings.NewReader(requestBody)
	client := &http.Client{}
	route, err := helpers.GetRoute(
		golang.RouteBlockPodsReorderBlockPodChecklists,
		*checklistParam.PodId,
		checklistParam.KeyId,
		*checklistParam.BlockId,
	)
	if err != nil {
		fmt.Println(err)
		return resChecklists.Checklists, err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resChecklists.Checklists, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resChecklists.Checklists, err
	}

	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resChecklists.Checklists, err
	}

	err = json.Unmarshal(body, &resChecklists)
	if err != nil {
		fmt.Println(err)
		return resChecklists.Checklists, err
	}
	return resChecklists.Checklists, nil
}
