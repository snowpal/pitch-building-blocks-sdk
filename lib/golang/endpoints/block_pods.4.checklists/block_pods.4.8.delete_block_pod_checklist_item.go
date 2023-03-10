package block_pods_4

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"fmt"
	"net/http"
)

func DeleteBlockPodChecklistItem(jwtToken string, checklistParam request.ChecklistIdParam) error {
	client := &http.Client{}
	route, err := helpers.GetRoute(
		golang.RouteBlockPodsDeleteBlockPodChecklistItem,
		checklistParam.PodId,
		checklistParam.ChecklistId,
		*checklistParam.ChecklistItemId,
		checklistParam.KeyId,
		checklistParam.BlockId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest(http.MethodDelete, route, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
