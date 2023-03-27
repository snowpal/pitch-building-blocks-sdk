package projectKeys

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/common"
	"development/go/recipes/lib/structs/response"
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
	route, err := helpers2.GetRoute(
		lib.RouteProjectKeysGetProjectLists,
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

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resProjectLists.ProjectLists, err
	}

	defer helpers2.CloseBody(res.Body)

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
