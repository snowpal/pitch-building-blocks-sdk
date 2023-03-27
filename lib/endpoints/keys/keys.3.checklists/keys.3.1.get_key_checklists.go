package keys

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/request"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetKeyChecklists(jwtToken string, checklistParam request.ChecklistIdParam) ([]response.Checklist, error) {
	resChecklists := response.Checklists{}
	route, err := helpers2.GetRoute(
		lib.RouteKeysGetKeyChecklists,
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

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resChecklists.Checklists, err
	}

	defer helpers2.CloseBody(res.Body)

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
