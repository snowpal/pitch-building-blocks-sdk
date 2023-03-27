package keyPods

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/request"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func AddKeyPodChecklist(
	jwtToken string,
	reqBody request.ChecklistReqBody,
	checklistParam request.ChecklistIdParam,
) (response.Checklist, error) {
	resChecklist := response.Checklist{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resChecklist, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteKeyPodsAddKeyPodChecklist,
		*checklistParam.PodId,
		checklistParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resChecklist, err
	}

	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resChecklist, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resChecklist, err
	}

	defer helpers2.CloseBody(res.Body)

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
