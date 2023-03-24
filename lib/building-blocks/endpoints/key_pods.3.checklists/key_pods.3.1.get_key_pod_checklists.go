package key_pods_3

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/request"
	"development/go/recipes/lib/building-blocks/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetKeyPodChecklists(jwtToken string, checklistParam request.ChecklistIdParam) ([]response.Checklist, error) {
	resChecklists := response.Checklists{}
	route, err := helpers.GetRoute(
		building_blocks.RouteKeyPodsGetKeyPodChecklists,
		*checklistParam.PodId,
		checklistParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resChecklists.Checklists, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resChecklists.Checklists, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
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
