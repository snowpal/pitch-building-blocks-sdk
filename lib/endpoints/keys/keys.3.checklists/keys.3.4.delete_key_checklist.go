package keys

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/request"
	"fmt"
	"net/http"
)

func DeleteKeyChecklist(jwtToken string, checklistParam request.ChecklistIdParam) error {
	route, err := helpers.GetRoute(
		lib.RouteKeysDeleteKeyChecklist,
		*checklistParam.ChecklistId,
		checklistParam.KeyId,
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

	_, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
