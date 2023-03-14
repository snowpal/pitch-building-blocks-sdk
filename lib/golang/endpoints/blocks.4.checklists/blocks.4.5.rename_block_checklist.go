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

func RenameBlockChecklist(
	jwtToken string,
	reqBody request.ChecklistReqBody,
	checklistParam request.ChecklistIdParam,
) (response.Checklist, error) {
	resChecklist := response.Checklist{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resChecklist, err
	}
	payload := strings.NewReader(requestBody)
	client := &http.Client{}
	route, err := helpers.GetRoute(
		golang.RouteBlocksRenameBlockChecklist,
		*checklistParam.BlockId,
		*checklistParam.ChecklistId,
		checklistParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resChecklist, err
	}
	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resChecklist, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resChecklist, err
	}

	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resChecklist, err
	}

	err = json.Unmarshal(body, &resChecklist)
	if err != nil {
		fmt.Println(err)
		return resChecklist, err
	}
	return resChecklist, nil
}