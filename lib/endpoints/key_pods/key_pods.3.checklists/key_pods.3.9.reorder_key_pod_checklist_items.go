package keyPods

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func ReorderKeyPodChecklistItems(
	jwtToken string,
	reqBody request.ReorderChecklistItemsReqBody,
	checklistParam request.ChecklistIdParam,
) ([]response.ChecklistItem, error) {
	resChecklistItems := response.ChecklistItems{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resChecklistItems.ChecklistItems, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteKeyPodsReorderKeyPodChecklistItems,
		*checklistParam.PodId,
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

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resChecklistItems.ChecklistItems, err
	}

	defer helpers2.CloseBody(res.Body)

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
