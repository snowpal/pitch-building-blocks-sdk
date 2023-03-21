package project_keys_2

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/common"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetProjectLists(
	jwtToken string,
	projectListParam common.ResourceIdParam,
) ([]response.ProjectList, error) {
	resProjectLists := response.ProjectLists{}
	route, err := helpers.GetRoute(
		golang.RouteProjectKeysGetProjectLists,
		projectListParam.BlockId,
		projectListParam.KeyId,
	)
	if err != nil {
		fmt.Println(err)
		return resProjectLists.ProjectLists, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resProjectLists.ProjectLists, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resProjectLists.ProjectLists, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resProjectLists.ProjectLists, err
	}

	err = json.Unmarshal(body, &resProjectLists)
	if err != nil {
		fmt.Println(err)
		return resProjectLists.ProjectLists, err
	}
	return resProjectLists.ProjectLists, nil
}
